package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id)

	quoteText := quote.Go()
	fmt.Printf("Random Quote: %s\n", quoteText)
}
