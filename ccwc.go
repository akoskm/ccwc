package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

/**
 * Parse program arguments and turn them into a map
 * program -c test -f -d best -l -g => map[-c:test -d:best -f: -g: -l:]
 *
 */
func parseParams(args []string) (map[string]string, error) {
	if len(args) == 0 {
		return nil, errors.New("No arguments provided.")
	}
	options := make(map[string]string)
	for i, arg := range args {
		// this is a new command starting with "-"
		if arg[0] == '-' {
			options[arg] = ""

			// let's see if it has a parameter
			if i+1 > len(args)-1 {
				continue
			}
			hasParam := args[i+1]
			if hasParam != "" && hasParam[0] != '-' {
				options[arg] = hasParam
			}
		}
	}
	return options, nil
}

func countBytes(filename string) {
	file, error := os.ReadFile(filename)
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(len(file), filename)
}

func countLines(filename string) {
	file, error := os.Open(filename)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	fmt.Println(lineCount, filename)
}

func countWords(filename string) {
	file, error := os.Open(filename)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	fmt.Println(wordCount, filename)
}

func countChars(filename string) {
	file, error := os.Open(filename)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	charCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		charCount += len(line)
	}

	fmt.Println(charCount, filename)
}

func main() {
	args := os.Args[1:]
	result, errors := parseParams(args)

	if errors != nil {
		fmt.Println(errors)
		return
	}

	if _, isPresent := result["-c"]; isPresent {
		countBytes(result["-c"])
	} else if _, isPresent := result["-l"]; isPresent {
		countLines(result["-l"])
	} else if _, isPresent := result["-w"]; isPresent {
		countWords(result["-w"])
	} else if _, isPresent := result["-m"]; isPresent {
		countChars(result["-m"])
	}
}
