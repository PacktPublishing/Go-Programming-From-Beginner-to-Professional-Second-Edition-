package main

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"File input", "Type in input"}

type model struct {
	cursor int
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("Select if you would like to work with file input or type in input:\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

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

// processFile processes either a file or user input, applies rot13 encoding, and writes to stdout.
func processFile(filename string) {
	var inputReader io.Reader

	// Check if a file path is provided
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	inputReader = file

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
	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok && m.choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.choice)
	}
	if m, ok := m.(model); ok && m.choice != "" && m.choice == "File input" {
		processFile("data.txt")
	}
	if m, ok := m.(model); ok && m.choice != "" && m.choice == "Type in input" {
		processStdin()
	}
}