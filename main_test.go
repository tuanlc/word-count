package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

func TestFoo(t *testing.T) {
	t.Run("should count correct byte", func(t *testing.T) {
		cmd := exec.Command("wc", "-c", "test.txt")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		expected := removeExtraSpaces(string(stdout))

		input := []string{"", "-c", "test.txt"}
		actual, err := foo(input, nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
	t.Run("should count correct line", func(t *testing.T) {
		cmd := exec.Command("wc", "-l", "test.txt")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		expected := removeExtraSpaces(string(stdout))

		input := []string{"", "-l", "test.txt"}
		actual, err := foo(input, nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
	t.Run("should count correct word", func(t *testing.T) {
		cmd := exec.Command("wc", "-w", "test.txt")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		expected := removeExtraSpaces(string(stdout))

		input := []string{"", "-w", "test.txt"}
		actual, err := foo(input, nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
	t.Run("should count correct characters", func(t *testing.T) {
		cmd := exec.Command("wc", "-m", "test.txt")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		expected := removeExtraSpaces(string(stdout))

		input := []string{"", "-m", "test.txt"}
		actual, err := foo(input, nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
	t.Run("should return outcome when there is no specific operation", func(t *testing.T) {
		cmd := exec.Command("wc", "test.txt")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		expected := removeExtraSpaces(string(stdout))

		input := []string{"", "test.txt"}
		actual, err := foo(input, nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
	t.Run("Standard input", func(t *testing.T) {
		t.Run("should return correct count byte ", func(t *testing.T) {
			cmd := exec.Command("bash", "-c", "cat test.txt | wc -c")
			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Errorf("Failed to exec wc command %s", err.Error())
			}

			expected := removeExtraSpaces(string(output))
			stdin, err := os.ReadFile("test.txt")

			if err != nil {
				t.Errorf("Failed to read file %s", err.Error())
			}

			input := []string{"", "-c"}
			actual, err := foo(input, stdin)

			if err != nil {
				t.Errorf("Failed to count %s", err.Error())
			}

			if expected != removeExtraSpaces(actual) {
				t.Errorf("expected '%v' but got '%v'", expected, removeExtraSpaces(actual))
			}
		})
		t.Run("should return correct count line ", func(t *testing.T) {
			cmd := exec.Command("bash", "-c", "cat test.txt | wc -l")
			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Errorf("Failed to exec wc command %s", err.Error())
			}

			expected := removeExtraSpaces(string(output))
			stdin, err := os.ReadFile("test.txt")

			if err != nil {
				t.Errorf("Failed to read file %s", err.Error())
			}

			input := []string{"", "-l"}
			actual, err := foo(input, stdin)

			if err != nil {
				t.Errorf("Failed to count %s", err.Error())
			}

			if expected != removeExtraSpaces(actual) {
				t.Errorf("expected '%v' but got '%v'", expected, removeExtraSpaces(actual))
			}
		})
		t.Run("should return correct count word ", func(t *testing.T) {
			cmd := exec.Command("bash", "-c", "cat test.txt | wc -w")
			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Errorf("Failed to exec wc command %s", err.Error())
			}

			expected := removeExtraSpaces(string(output))
			stdin, err := os.ReadFile("test.txt")

			if err != nil {
				t.Errorf("Failed to read file %s", err.Error())
			}

			input := []string{"", "-w"}
			actual, err := foo(input, stdin)

			if err != nil {
				t.Errorf("Failed to count %s", err.Error())
			}

			if expected != removeExtraSpaces(actual) {
				t.Errorf("expected '%v' but got '%v'", expected, removeExtraSpaces(actual))
			}
		})
		t.Run("should return correct count character ", func(t *testing.T) {
			cmd := exec.Command("bash", "-c", "cat test.txt | wc -m")
			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Errorf("Failed to exec wc command %s", err.Error())
			}

			expected := removeExtraSpaces(string(output))
			stdin, err := os.ReadFile("test.txt")

			if err != nil {
				t.Errorf("Failed to read file %s", err.Error())
			}

			input := []string{"", "-m"}
			actual, err := foo(input, stdin)

			if err != nil {
				t.Errorf("Failed to count %s", err.Error())
			}

			if expected != removeExtraSpaces(actual) {
				t.Errorf("expected '%v' but got '%v'", expected, removeExtraSpaces(actual))
			}
		})
		t.Run("should return correct count when there is no specific operation", func(t *testing.T) {
			cmd := exec.Command("bash", "-c", "cat test.txt | wc")
			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Errorf("Failed to exec wc command %s", err.Error())
			}

			expected := removeExtraSpaces(string(output))
			stdin, err := os.ReadFile("test.txt")

			if err != nil {
				t.Errorf("Failed to read file %s", err.Error())
			}

			input := []string{""}
			actual, err := foo(input, stdin)

			if err != nil {
				t.Errorf("Failed to count %s", err.Error())
			}

			if expected != removeExtraSpaces(actual) {
				t.Errorf("expected '%v' but got '%v'", expected, removeExtraSpaces(actual))
			}
		})
	})
}

// removeExtraSpaces removes extra spaces between words in a string
func removeExtraSpaces(input string) string {
	spaceRegex := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(spaceRegex.ReplaceAllString(input, " "))
}
