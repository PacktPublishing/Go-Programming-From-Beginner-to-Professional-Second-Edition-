package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// rot13 applies rot13 encoding to a given string.
func rot13(s string) string {
	result := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		char := s[i]

		switch {
		case char >= 'a' && char <= 'z':
			result[i] = 'a' + (char-'a'+13)%26
		case char >= 'A' && char <= 'Z':
			result[i] = 'A' + (char-'A'+13)%26
		default:
			result[i] = char
		}
	}

	return string(result)
}

// processStdin reads data from stdin, applies rot13 encoding, and writes to stdout.
func processStdin() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading stdin:", err)
			return
		}
		encoded := rot13(input)
		fmt.Print(encoded)
	}
}

// processFileOrInput processes either a file or user input, applies rot13 encoding, and writes to stdout.
func processFileOrInput() {
	var inputReader io.Reader

	// Check if a file path is provided
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		inputReader = file
	} else {
		// No file provided, read user input
		fmt.Print("Enter text: ")
		inputReader = os.Stdin
	}

	// Process input and apply rot13 encoding
	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		// Apply rot13 encoding to the input line
		encoded := rot13(scanner.Text())
		fmt.Println(encoded)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func main() {
	// Check if data is available on stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data available on stdin, process it
		processStdin()
	} else {
		// No data on stdin, process file or user input
		processFileOrInput()
	}
}
