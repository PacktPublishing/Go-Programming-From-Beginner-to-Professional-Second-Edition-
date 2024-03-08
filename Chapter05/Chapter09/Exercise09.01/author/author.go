package author

import (
	"fmt"
)

// Author represents an author of a book.
type Author struct {
	Name    string
	Contact string
}

// NewAuthor creates a new Author instance.
func NewAuthor(name, contact string) *Author {
	return &Author{Name: name, Contact: contact}
}

// WriteChapter writes a book chapter by the author.
func (a *Author) WriteChapter(chapterTitle string, content string) {
	fmt.Printf("Author %s is writing a chapter titled '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

// ReviewChapter allows the author to review and edit a chapter.
func (a *Author) ReviewChapter(chapterTitle string, content string) {
	fmt.Printf("Author %s is reviewing a chapter titled '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

// FinalizeChapter marks the chapter as finalized by the author.
func (a *Author) FinalizeChapter(chapterTitle string) {
	fmt.Printf("Author %s has finalized the chapter titled '%s'.\n", a.Name, chapterTitle)
}
