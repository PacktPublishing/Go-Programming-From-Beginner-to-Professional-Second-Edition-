package main

/*
import (
	"fmt"
	"os"
	"time"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the state of the TUI.
type Model struct {
	SelectionMode string // "text" or "file"
	Lines         []string
	CurrentLine   string
	FileName      string
	Prompt        string
	Error         string
	FileSelected  bool
}

// Message represents the user's interaction with the TUI.
type Message string

const (
	SelectText Message = "Select Text"
	SelectFile Message = "Select File"
	Finish     Message = "Finish"
)

// Init initializes the model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update processes messages received by the TUI.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+d":
			return m, tea.Quit
		}
	case tea.MouseMsg:
		if msg.Type == tea.MouseWheelDown {
			return m, nil
		}
	}

	switch msg {
	case SelectText:
		m.SelectionMode = "T"
		m.Prompt = "Enter text (Ctrl+D to finish):"
		return m, nil

	case SelectFile:
		m.SelectionMode = "F"
		m.Prompt = "Enter file name:"
		return m, nil

	case Finish:
		return m, tea.Quit
	}

	return m, tea.Tick(time.Second, func(time.Time) tea.Msg {
		return nil
	})
}

// View renders the TUI components.
func (m Model) View() string {
	var view string

	if m.SelectionMode == "" {
		view = "Choose input mode:\n[T] Text\n[F] File\n"
	} else if m.SelectionMode == "T" {
		view = fmt.Sprintf("%s\n%s", m.Prompt, formatLines(m.Lines))
	} else if m.SelectionMode == "F" {
		view = fmt.Sprintf("%s %s", m.Prompt, m.FileName)
	}

	if m.Error != "" {
		view += fmt.Sprintf("\n\nError: %s", m.Error)
	}

	view += fmt.Sprintf("\n\n%s", m.Prompt)

	return view
}

func formatLines(lines []string) string {
	result := ""
	for _, line := range lines {
		result += rot13(line) + "\n"
	}
	return result
}

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

func (m Model) Cursor() (int, int) {
	x, y := 0, utf8.RuneCountInString(m.Prompt+"\n")

	if m.SelectionMode == "text" {
		x = utf8.RuneCountInString(m.Prompt + "\n" + formatLines(m.Lines))
	}

	if m.SelectionMode == "file" {
		x = utf8.RuneCountInString(m.Prompt + " " + m.FileName)
	}

	return x, y
}

func main() {
	p := tea.NewProgram(Model{
		Prompt: "",
	})
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting TUI: %v\n", err)
		os.Exit(1)
	}
}
*/
