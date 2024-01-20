package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func countAll(filename string) map[string]int {
	file, error := os.Open(filename)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	fileBytes := 0
	charCount := 0
	wordCount := 0
	lineCount := 0
	inWord := false
	for {
		char, size, err := reader.ReadRune()

		fileBytes += size

		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		charCount++
		if unicode.IsSpace(char) {
			if char == '\n' {
				lineCount++
			}
			if inWord {
				wordCount++
			}
			inWord = false
		} else {
			inWord = true
		}
	}
	result := map[string]int{
		"-c": fileBytes,
		"-l": lineCount,
		"-w": wordCount,
		"-m": charCount,
	}
	return result
}

/**
 * Prefix a number with spaces until it hits @spacing length
 */
func formatNumber(num int, spacing int) string {
	strNum := strconv.Itoa(num)
	numLength := len(strNum)
	result := ""

	if numLength >= spacing {
		return result
	}

	for _, char := range strNum {
		result += string(char)
	}

	for i := 0; i < spacing-numLength; i++ {
		result = " " + result
	}
	return result
}

func isFlagPresent(flag string, args []string) bool {
	for i := 0; i < len(args); i++ {
		if args[i] == flag {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]

	filename := args[len(args)-1]

	parsed := countAll(filename)

	display := filename
	flags := []string{"-c", "-l", "-w", "-m"}
	for i := 0; i < len(flags); i++ {
		currArg := flags[i]
		if isFlagPresent(currArg, args) {
			formattedResult := formatNumber(parsed[currArg], 8)
			display = formattedResult + " " + display
		}
	}

	fmt.Println(display)
}
