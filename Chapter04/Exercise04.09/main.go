package main

import (
	"fmt"
	"os"
)

func getPassedArgs() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func getLocales(extraLocales []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocales...)
	return locales
}

func main() {
	locales := getLocales(getPassedArgs())
	fmt.Println("Locales to use:", locales)
}
