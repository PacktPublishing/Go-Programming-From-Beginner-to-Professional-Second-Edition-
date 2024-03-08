package main

import (
	"flag"
	"fmt"
)

var (
	nameFlag  = flag.String("name", "Sam", "Name of the person to say hello to")
	quietFlag = flag.Bool("quiet", false, "Toggle to be quiet when saying hello")
)

func main() {
	flag.Parse()

	if !*quietFlag {
		greeting := fmt.Sprintf("Hello, %s! Welcome to the command line.", *nameFlag)
		fmt.Println(greeting)
	}
}
