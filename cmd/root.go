package cmd

import (
	"fmt"
	"os"

	"github.com/deependujha/litracer/os_utils"
	"github.com/deependujha/litracer/parser"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "litracer",
	Short: "A tool for converting Lightning AI's LitData logs to Chrome-compatible trace files",
	Long:  `A tool for converting Lightning AI's LitData logs to Chrome-compatible trace files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Entry point for the command
		if len(args) != 1 {
			fmt.Println("only one argument is required `litracer <log_file_path>`. But received", len(args), "arguments: ", args)
			os.Exit(1)
		}

		fmt.Println("Arguments:", args)
		file_path := args[0]
		if !os_utils.DoesFileExist(file_path) {
			fmt.Println("File does not exist:", file_path)
			os.Exit(1)
		}
		fmt.Println("File exists:", file_path)
		parser.ParseFile(file_path, 1)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
