package bubbletea

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/deependujha/litracer/litparser"
	"github.com/deependujha/litracer/os_utils"
)

func Run(log_file_path string, numWorkers int, outputFilepath string) {
	// Print the details of the file and the number of workers
	worker_emoji := workerEmoji()

	// Print the emoji
	file_log := fmt.Sprintf("\t📝 Tracing log file: %s", log_file_path)
	worker_log := fmt.Sprintf("\t%s Number of Workers: %d\n", worker_emoji, numWorkers)

	fmt.Println(YELLOW + file_log + RESET)
	fmt.Printf(MAGENTA + worker_log + RESET)
	// get number of lines in the file
	helperCh := make(chan os_utils.NumberOfLinesAndError)
	go os_utils.GetNumberOfLines(log_file_path, helperCh)

	numberOfLines := 0
loop:
	for {
		// sleep for 0.1 second and check if channel has any data, until true, repeat
		select {
		case result := <-helperCh:
			_ = result
			if result.Error != nil {
				fmt.Printf("Error getting number of lines: %s\n", result.Error)
				os.Exit(1)
			}
			numberOfLines = result.NumberOfLines + 1 // add 1 to include the last line
			break loop

		default:
			stopped := KeepSpinning(1500, "Getting number of lines in the log file") // keep spinning for 2 ms
			if stopped {
				break loop
			}

		}
	}
	fmt.Print("\r\033[K")
	number_of_lines_log := fmt.Sprintf("\t📄 Number of lines in the log file: %d\n", numberOfLines)
	fmt.Println(GREEN + number_of_lines_log + RESET)
	start(log_file_path, numWorkers, outputFilepath, numberOfLines)

	fmt.Println("✅ Parsing completed.")
	fmt.Printf("⚡️ Use " + RED + "chrome://tracing" + RESET + " or " + RED + "ui.perfetto.dev " + RESET + "to view the trace file: " + YELLOW + outputFilepath + RESET + "\n")
}

func start(log_file_path string, numWorkers int, outputFilepath string, numberOfLines int) {
	// start progress bar
	parsedLinesChan := make(chan int) // channel to send number of lines parsed

	go litparser.ParseFile(log_file_path, numWorkers, outputFilepath, parsedLinesChan)
	StartProgressBar(numberOfLines, parsedLinesChan)

}

func workerEmoji() string {
	xx := "\\U0001f477" // this is 👷

	// Convert to hex string
	h := strings.ReplaceAll(xx, "\\U", "0x")
	// Hex to Int
	i, _ := strconv.ParseInt(h, 0, 32) // limit to 32 bits
	// Convert to rune, then to string
	str := string(rune(i))
	return str
}
