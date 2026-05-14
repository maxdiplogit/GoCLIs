package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const KILOBYTE int = 1024 // 1KB = 1024 Bytes
const MAX_FOUND_LINES int = 10

// This tool will look up for substring matches in a file
// At whatever line it finds a match for the provided substring, it will return the match
func main() {
	userWorkingDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./mygrep <substring> filename")
		os.Exit(1)
	}

	substring, filename := os.Args[1], os.Args[2]

	filePath := filepath.Join(userWorkingDir, filename)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*KILOBYTE), 128*KILOBYTE)

	lines := make([]string, 0, MAX_FOUND_LINES)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, substring) {
			lines = append(lines, line)
			continue
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
