package os_utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// DoesFileExist checks if a file exists and is not a directory
func DoesFileExist(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ReadFile reads the entire content of a file and returns it as a string
func ReadFile(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(data), nil
}

// ReadFileLineByLine reads a file line by line and returns a slice of strings
func ReadFileLineByLine(filepath string, ch chan string) error {
	defer close(ch)

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

// WriteToFile writes data to a file, overwriting it if it exists
func WriteToFile(filepath, data string) error {
	return os.WriteFile(filepath, []byte(data), 0644)
}

// AppendToFile appends data to an existing file or creates it if it doesn't exist
func AppendToFile(filepath, data string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for appending: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return fmt.Errorf("failed to append to file: %w", err)
	}

	return nil
}

// uses `wc -l` to count the number of lines in a file
func GetNumberOfLines(filepath string, ch chan NumberOfLinesAndError) {
	cmd := exec.Command("wc", "-l", filepath)
	output, err := cmd.Output()

	if err != nil {
		ch <- NumberOfLinesAndError{0, err}
		return
	}
	// expected output: "  1234 filename.txt"
	// split by space and take the first part and convert to int
	// trim the output to remove leading and trailing spaces
	// and split by space
	trimmed_output := strings.TrimSpace(string(output))

	splitted_trimmed_output := strings.Split(trimmed_output, " ")
	var numLines int
	if len(splitted_trimmed_output) > 1 {
		numLines, err = strconv.Atoi(strings.TrimSpace(splitted_trimmed_output[0]))
		if err != nil {
			ch <- NumberOfLinesAndError{0, err}
			return
		}
		ch <- NumberOfLinesAndError{numLines, nil}
		return
	} else {
		ch <- NumberOfLinesAndError{
			0,
			fmt.Errorf("unexpected output format for `wc -l file.txt`. Expected: '  1234 filename.txt', got: '%s'", trimmed_output),
		}
		return
	}
}
