package trace_writer

import (
	"fmt"
	"os"
)

func TraceWriter(filepath string, sinkLimit int, content chan JsonContent) {

	buffered_map := make(map[int]string)
	current_line := 1

	firstJsonDone := false // first json should not have a comma before it

	output_chan := make(chan string)
	defer close(output_chan)

	go TraceEventSink(filepath, sinkLimit, output_chan)

	for cnt := range content {
		line_no := cnt.LineNo
		txt := cnt.Content
		if line_no == current_line {
			if txt != "" {
				if firstJsonDone {
					output_chan <- "," + txt
				} else {
					output_chan <- txt
					firstJsonDone = true
				}
			}
			current_line++

			// keep looking in the buffered map until we don't find the next one
			for {
				if _, ok := buffered_map[current_line]; ok {
					txt = buffered_map[current_line]
					delete(buffered_map, current_line)
					if txt != "" {
						if firstJsonDone {
							output_chan <- "," + txt
						} else {
							output_chan <- txt
							firstJsonDone = true
						}
					}
					current_line++
				} else {
					break
				}
			}

		} else {
			buffered_map[line_no] = txt
		}
	}
}

func TraceEventSink(filepath string, sinkLimit int, ch <-chan string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("failed to open file for appending: %w", err))
	}
	defer file.Close()

	content_array := make([]string, sinkLimit)

	idx := 0
	for content := range ch {
		content_array[idx] = content
		idx++
		if idx >= sinkLimit {
			all_content := ""
			for i := 0; i < len(content_array); i++ {
				all_content += content_array[i]
				content_array[i] = "" // Clear the content array
			}
			_, err := file.WriteString(all_content)
			if err != nil {
				panic(fmt.Errorf("failed to write to file: %w", err))
			}
			// Reset the index and clear the content array
			idx = 0
		}
	}

	// when channel closes, some content may be left in the array
	// let's say we have 15 lines to write with sinkLimit = 10
	// we will have 5 lines left in the array, even if the channel closes
	if idx > 0 {
		all_content := ""
		for i := 0; i < idx; i++ {
			all_content += content_array[i]
			content_array[i] = "" // Clear the content array
		}
		_, err := file.WriteString(all_content)
		if err != nil {
			panic(fmt.Errorf("failed to write to file: %w", err))
		}
	}
}
