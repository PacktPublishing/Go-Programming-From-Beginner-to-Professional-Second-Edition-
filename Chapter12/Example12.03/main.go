package main

import (
	"fmt"
	"time"
)

func main() {
	deadline_seconds := time.Duration((600 * 10) * time.Millisecond)

	Start := time.Now()

	fmt.Println("Deadline for the transaction is ", deadline_seconds)

	fmt.Println("The transaction has started at: ", Start)

	sum := 0

	for i := 1; i < 25000000000; i++ {

		sum += i

	}

	End := time.Now()

	Duration := End.Sub(Start)

	TransactionTime := time.Duration(Duration.Nanoseconds()) * time.Nanosecond

	fmt.Println("The transaction has completed at: ", End, Duration)

	if TransactionTime <= deadline_seconds {

		fmt.Println("Performance is OK transaction completed in", TransactionTime)

	} else {

		fmt.Println("Performance problem, transaction completed in", TransactionTime, "second(s)!")

	}
}
