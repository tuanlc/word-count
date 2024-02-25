package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outcome, err := foo(os.Args, stdin)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(outcome)
}

func foo(args []string, stdin []byte) (outcome string, err error) {
	operation, filePath := getParameters(args)
	content, err := getContent(filePath, stdin)

	if err != nil {
		return outcome, err
	}

	switch operation {
	case "-c":
		{
			count := byteCount(content)
			outcome = fmt.Sprintf("%v %s", count, filePath)
			break
		}
	case "-l":
		{
			count := lineCount(content)
			outcome = fmt.Sprintf("%v %s", count, filePath)
			break
		}
	case "-w":
		{
			count := wordCount(content)
			outcome = fmt.Sprintf("%v %s", count, filePath)
			break
		}
	case "-m":
		{
			count := characterCount(content)
			outcome = fmt.Sprintf("%v %s", count, filePath)
			break
		}
	case "":
		{
			byteCount := byteCount(content)
			lineCount := lineCount(content)
			wordCount := wordCount(content)
			outcome = fmt.Sprintf("%v %v %v %s", lineCount, wordCount, byteCount, filePath)
			break
		}
	default:
		{
			return outcome, fmt.Errorf("Invalid operation %s", operation)
		}
	}

	return outcome, nil
}

func getContent(filePath string, stdin []byte) (content []byte, err error) {
	if filePath != "" {
		return os.ReadFile(filePath)
	}

	return stdin, nil
}

func getParameters(args []string) (operation string, filePath string) {
	filePath = ""

	if len(args) == 2 {
		firstArg := args[1]

		switch firstArg {
		case "-c", "-l", "-w", "-m":
			operation = firstArg
			break
		default:
			filePath = firstArg
		}
	}

	if len(args) >= 3 {
		operation = args[1]
		filePath = args[2]
	}

	return operation, filePath
}

func byteCount(fileContent []byte) int {
	return len(fileContent)
}

func lineCount(fileContent []byte) int {
	lineSep := []byte{'\n'}
	return bytes.Count(fileContent, lineSep)
}

func wordCount(fileContent []byte) int {
	words := strings.Fields(string(fileContent))
	return len(words)
}

func characterCount(fileContent []byte) int {
	return utf8.RuneCountInString(string(fileContent))
}
