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

	operation, err := safeArrayAccess(args, 1)

	if err != nil {
		fmt.Println("Missing operation. Available operations: -c, -l, -w, -m")
		os.Exit(1)
	}

	if operation != "-c" && operation != "-l" && operation != "-w" && operation != "-m" {
		fmt.Println("Invalid operation. Available operations: -c, -l, -w, -m")
		os.Exit(1)
	}

	file, err := safeArrayAccess(args, 2)

	if err != nil {
		fmt.Println("Missing target file")
		os.Exit(1)
	}

	if operation == "-c" {
		count, err := byteCount(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%v %s \n", count, file)
	}

	if operation == "-l" {
		count, err := lineCount(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%v %s \n", count, file)
	}

	if operation == "-w" {
		count, err := wordCount(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%v %s \n", count, file)
	}

	if operation == "-m" {
		count, err := characterCount(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%v %s \n", count, file)
	}
}

func byteCount(file string) (count int, err error) {
	content, readError := os.ReadFile(file)
	if readError != nil {
		err = readError
	}

	count = len(content)

	return count, err
}

func lineCount(file string) (count int, err error) {
	content, readError := os.ReadFile(file)
	if readError != nil {
		err = readError
	}

	lineSep := []byte{'\n'}
	count = bytes.Count(content, lineSep)

	return count, err
}

func wordCount(file string) (count int, err error) {
	content, readError := os.ReadFile(file)
	if readError != nil {
		err = readError
	}

	words := strings.Fields(string(content))
	count = len(words)

	return count, err
}

func characterCount(file string) (count int, err error) {
	content, readError := os.ReadFile(file)
	if readError != nil {
		err = readError
	}

	count = utf8.RuneCountInString(string(content))

	return count, err
}
