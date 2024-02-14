package main

import "fmt"

func main() {
	finished := make(chan bool)
	names := []string{"Packt"}

	go func() {
		names = append(names, "Electric")
		names = append(names, "Boogaloo")
		finished <- true
	}()
	// <-finished

	for _, name := range names {
		fmt.Println(name)
	}
	<-finished
}
