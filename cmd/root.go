package cmd

import (
	"fmt"
	"os"

	"github.com/deependujha/litracer/bubbletea"
	"github.com/deependujha/litracer/os_utils"

	// "github.com/deependujha/litracer/parser"
	"github.com/spf13/cobra"
)

var outputFilepath string
var numWorkers int
var sinkLimit int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "litracer",
	Short: "A tool for converting Lightning AI's LitData logs to Chrome-compatible trace files",
	Long:  `A tool for converting Lightning AI's LitData logs to Chrome-compatible trace files.`,
	Args:  cobra.ExactArgs(1), // Enforce exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		// Entry point for the command
		if len(args) != 1 {
			fmt.Fprint(os.Stderr, bubbletea.RED+"Error: please provide a path to the input log file.\n\n"+bubbletea.RESET)
			fmt.Fprint(os.Stderr, bubbletea.YELLOW+"Usage: litracer [flags] <log_file_path>\n\n"+bubbletea.RESET)
			fmt.Fprint(os.Stderr, bubbletea.GREEN+"Help:\n"+bubbletea.RESET)
			cmd.Help()
			os.Exit(1)
		}

		// if numWorkers < 1 || numWorkers > 1 {
		// 	fmt.Println("Only 1 worker is supported at the moment")
		// 	os.Exit(1)
		// }

		log_file_path := args[0]

		if !os_utils.DoesFileExist(log_file_path) {
			fmt.Println("File does not exist:", log_file_path)
			os.Exit(1)
		}
		fmt.Println() // Print a new line for better readability
		bubbletea.Run(log_file_path, numWorkers, sinkLimit, outputFilepath)
	},
}

func init() {
	// Define flags
	rootCmd.Flags().StringVarP(&outputFilepath, "output", "o", "litdata_trace.json", "Path to the output trace file")
	rootCmd.Flags().IntVarP(&numWorkers, "workers", "w", 1, "Number of worker goroutines to use for parsing")
	rootCmd.Flags().IntVarP(&sinkLimit, "sink", "s", 100, "Sink limit: number of lines to write at once")
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
