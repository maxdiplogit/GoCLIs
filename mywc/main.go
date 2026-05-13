package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	// "flag"
)

const KILOBYTE uint16 = 1024

// This tool reimplements UNIX's wc binary in Go
// It reads either from the stdin/command-line-args or from a file
// Currently it has only two flags -w (count only words) and -l (count only lines)

func main() {
	// cli_args := os.Args

	// fmt.Printf("CLI Args: %#v\n", cli_args)

	// for i := 0; i < len(cli_args); i++ {
	// 	fmt.Printf("Arg %d: %s\n", i, cli_args[i])
	// 	fmt.Printf("Arg length %d: %d\n", i, len(cli_args[i]))
	// 	for j := 0; j < len(cli_args[i]); j++ {
	// 		fmt.Printf("Byte: %#v\n", cli_args[i][j])
	// 	}
	// }

	home, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit does not run deferred functions.
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./mywc <filename>")
		// os.Exit does not run deferred functions.
		os.Exit(1)
	}

	filePath := filepath.Join(home, os.Args[1])

	fmt.Printf("FilePath: %s\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Buffer(make([]byte, 64*int(KILOBYTE)), 128*int(KILOBYTE))

	var lines, words, letters int = 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println("Line ", lines, ": ", line)

		wordsFromString := strings.Split(line, " ")
		// fmt.Printf("wordsFromString: %#v\n", wordsFromString)

		for _, word := range wordsFromString {
			// fmt.Printf("word: %#v\n", word)

			for range word {
				// fmt.Printf("letter: %#v\n", letter)
				letters++
			}
		}

		words += len(wordsFromString)

		lines++
	}

	fmt.Println("Letters: ", letters)
	fmt.Println("Words: ", words)
	fmt.Println("Lines: ", lines)
}
