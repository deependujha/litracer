package os_utils

import (
	"bufio"
	"fmt"
	"os"
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
