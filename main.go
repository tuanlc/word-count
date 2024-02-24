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
	operation, filePath := getParameters()
	content, err := getContent(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch operation {
	case "-c":
		{
			count := byteCount(content)
			fmt.Printf("%v %s \n", count, filePath)
			break
		}
	case "-l":
		{
			count := lineCount(content)
			fmt.Printf("%v %s \n", count, filePath)
			break
		}
	case "-w":
		{
			count := wordCount(content)
			fmt.Printf("%v %s \n", count, filePath)
			break
		}
	case "-m":
		{
			count := characterCount(content)
			fmt.Printf("%v %s \n", count, filePath)
			break
		}
	case "":
		{
			byteCount := byteCount(content)
			lineCount := lineCount(content)
			wordCount := wordCount(content)
			fmt.Printf("Byte   Line   Word   File \n")
			fmt.Printf("%v   %v   %v %s \n", byteCount, lineCount, wordCount, filePath)
			break
		}
	default:
		{
			fmt.Printf("Invalid operation")
			os.Exit(1)
		}
	}
}

func getContent(filePath string) (content []byte, err error) {
	if filePath != "" {
		return os.ReadFile(filePath)
	}

	data, readErr := io.ReadAll(os.Stdin)
	if readErr != nil {
		return content, readErr
	}

	return data, nil
}

func getParameters() (operation string, filePath string) {
	args := os.Args
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
