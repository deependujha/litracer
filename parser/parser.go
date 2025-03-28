package parser

import (
	"fmt"
	"strings"
	"sync"

	"github.com/deependujha/litracer/os_utils"
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
func worker(id int, lines <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	_ = id
	for line := range lines {
		_ = ParseLine(line)
	}
}

// ParseFile
// This function is used to parse the file.
// It reads the file line by line and distributes the work to the workers.
func ParseFile(filepath string, numWorkers int) {
	linesChan := make(chan string, numWorkers)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, linesChan, &wg)
	}

	if err := os_utils.ReadFileLineByLine(filepath, linesChan); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	close(linesChan) // No more lines to send
	wg.Wait()        // Wait for workers to finish

}
