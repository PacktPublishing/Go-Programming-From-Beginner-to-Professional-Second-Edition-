package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id)
}
