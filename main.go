package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var (
	path string
)

var availableOperations = [2]string{"-c", "-l"}

func safeArrayAccess(arr []string, index int) (value string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic: %v", r)
		}
	}()

	value = arr[index]
	return value, nil
}

func main() {
	args := os.Args
	filePath := ""
	operation := ""

	if len(args) == 1 {
		fmt.Println("Missing target file")
	}

	if len(args) == 2 {
		file, err := safeArrayAccess(args, 1)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		filePath = file
	}

	if len(args) >= 3 {
		operationInput, err := safeArrayAccess(args, 1)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		operation = operationInput

		file, err := safeArrayAccess(args, 2)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		filePath = file
	}

	if filePath == "" {
		fmt.Println("Missing target file")
		os.Exit(1)
	}

	content, err := os.ReadFile(filePath)
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
		}
	default:
		{
			fmt.Printf("Invalid operation")
			os.Exit(1)
		}
	}
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
