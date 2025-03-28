package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "litracer",
	Short: "A tool for converting Lightning AI's LitData logs to Chrome-compatible trace files",
	Long: `litracer is a command-line utility designed to process log files generated
by Lightning AI's LitData framework and converts them into json trace files that
can be used to visualize the trace using Chrome tracing tools.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Entry point for the command
		fmt.Println("litracer: processing log files...")

		// Print received arguments
		if len(args) > 0 {
			fmt.Println("Arguments:", args)
		} else {
			fmt.Println("No arguments provided.")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
