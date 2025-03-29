package parser

import (
	"fmt"
	"strings"
	"sync"

	"github.com/deependujha/litracer/os_utils"
	"github.com/deependujha/litracer/reflection_utils"
	"github.com/deependujha/litracer/trace_event"
)

// ParseLine
// Each line of the log file is a key-value pair separated by a semicolon.
// e.g. "key1:value1;key2:value2;key3:value3"
//
// This function parses the line and returns a map of the key-value pairs.
func ParseLine(line string) map[string]string {
	result := make(map[string]string)
	pairs := strings.Split(line, ";")

	for _, pair := range pairs {
		kv := strings.SplitN(pair, ":", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			result[key] = value
		}
	}

	return result
}

// worker
// This function is used to parse the lines in parallel.
// It reads from the channel and parses the line.
func worker(worker_id int, lines <-chan string, wg *sync.WaitGroup, output_file string) {
	defer wg.Done()
	_ = worker_id
	tr := trace_event.TraceEvent{}
	first_line := true
	for line := range lines {
		parsed_line := ParseLine(line)
		// content := fmt.Sprintf("worker_id: %d; %v", worker_id, parsed_line)
		// fmt.Println(content)

		err := reflection_utils.MapToStruct(parsed_line, &tr)
		if err != nil {
			fmt.Println("45: Error parsing line:", err)
			continue
		}
		json_data, err := tr.ToJSON()
		if err != nil {
			fmt.Println("45: Error parsing line:", err)
			continue
		}
		if first_line {
			os_utils.AppendToFile(output_file, json_data)
			first_line = false
		} else {
			os_utils.AppendToFile(output_file, ","+json_data)
		}
	}
}

// ParseFile
// This function is used to parse the file.
// It reads the file line by line and distributes the work to the workers.
func ParseFile(filepath string, numWorkers int, output_file string) {
	linesChan := make(chan string, numWorkers)
	// defer close(linesChan)

	var wg sync.WaitGroup

	os_utils.WriteToFile(output_file, "{\"traceEvents\":[")

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, linesChan, &wg, output_file)
	}

	if err := os_utils.ReadFileLineByLine(filepath, linesChan); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	wg.Wait() // Wait for workers to finish

	os_utils.AppendToFile(output_file, "]}")

}
