package litparser

import (
	"fmt"
	"strings"
	"sync"

	"github.com/deependujha/litracer/os_utils"
	"github.com/deependujha/litracer/reflection_utils"
	"github.com/deependujha/litracer/trace_event"
	"github.com/deependujha/litracer/trace_writer"
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
			key := strings.ToLower(strings.TrimSpace(kv[0]))
			value := strings.TrimSpace(kv[1])
			result[key] = value
		}
	}

	return result
}

// worker
// This function is used to parse the lines in parallel.
// It reads from the channel and parses the line.
func worker(worker_id int, lines <-chan trace_writer.JsonContent, wg *sync.WaitGroup, parsedLinesChan chan int, parsedJsonChan chan trace_writer.JsonContent) {
	defer wg.Done()
	_ = worker_id
	tr := trace_event.TraceEvent{}

	for line := range lines {
		parsed_line := ParseLine(line.Content)
		parsedLinesChan <- 1 // Send the number of lines parsed to the channel
		// If the 'ph' key is missing, skip processing this line.
		if _, ok := parsed_line["ph"]; !ok {
			parsedJsonChan <- trace_writer.JsonContent{LineNo: line.LineNo, Content: ""}
			continue
		}

		err := reflection_utils.MapToStruct(parsed_line, &tr)
		if err != nil {
			fmt.Println("45: Error parsing line:", err)
			continue
		}
		json_data, err := tr.ToJSON()
		if err != nil {
			fmt.Println("Error parsing line to json:", err)
			continue
		}
		parsedJsonChan <- trace_writer.JsonContent{LineNo: line.LineNo, Content: json_data}
	}
}

// ParseFile
// This function is used to parse the file.
// It reads the file line by line and distributes the work to the workers.
func ParseFile(filepath string, numWorkers int, sinkLimit int, output_file string, parsedLinesCountChan chan int) {
	defer close(parsedLinesCountChan)

	linesChan := make(chan trace_writer.JsonContent)
	parsedJsonChan := make(chan trace_writer.JsonContent)

	var wg sync.WaitGroup

	// delete the output file if it exists
	os_utils.DeleteFile(output_file)
	os_utils.WriteToFile(output_file, "{\"traceEvents\":[")
	go trace_writer.TraceWriter(output_file, sinkLimit, parsedJsonChan)

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, linesChan, &wg, parsedLinesCountChan, parsedJsonChan)
	}

	if err := os_utils.ReadFileLineByLine(filepath, linesChan); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	wg.Wait() // Wait for workers to finish
	close(parsedJsonChan)

	os_utils.AppendToFile(output_file, "]}")

}
