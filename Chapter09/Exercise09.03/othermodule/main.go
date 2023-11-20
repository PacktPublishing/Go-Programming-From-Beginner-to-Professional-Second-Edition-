package main

import (
	"fmt"

	"github.com/sicoyle/printer"
)

func main() {
	msg := printer.PrintNewUUID()
	fmt.Println(msg)
}
