package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	nyTime, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	laTime, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The local current time is:", current.Format(time.ANSIC))
	fmt.Println("The time in New York is: ", current.In(nyTime).Format(time.ANSIC))
	fmt.Println("The time in Los Angeles is: ", current.In(laTime).Format(time.ANSIC))
}
