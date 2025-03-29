package trace_writer

import (
	"fmt"
	"os"

	"github.com/deependujha/litracer/os_utils"
	"github.com/deependujha/litracer/parser"
)

func TraceWriter(filepath string, content string) {
	os_utils.AppendToFile(filepath, content)
}

func TraceEventSink(filepath string, ch <-chan parser.SinkDType, numWorkers int) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("failed to open file for appending: %w", err))
	}
	defer file.Close()

	content_array := make([]string, numWorkers)

	for content := range ch {
		worker_id := content.WorkerID
		content_array[worker_id] = content.Content
	}

}
