package main

import (
	"fmt"
	"time"
)

func main() {
	Start := time.Now()

	fmt.Println("The script started at: ", Start)

	sum := 0

	for i := 1; i < 10000000000; i++ {

		sum += i

	}

	End := time.Now()

	Duration := End.Sub(Start)

	fmt.Println("The script completed at: ", End)

	fmt.Println("The task took", Duration.Hours(), "hour(s) to complete!")

	fmt.Println("The task took", Duration.Minutes(), "minutes(s) to complete!")

	fmt.Println("The task took", Duration.Seconds(), "seconds(s) to complete!")

	fmt.Println("The task took", Duration.Nanoseconds(), "nanosecond(s) to complete!")
}
