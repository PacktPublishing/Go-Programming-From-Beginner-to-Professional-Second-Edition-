package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println("The current time is:", current.Format(time.ANSIC))
	// 6 hours, 6 minutes, 6 seconds = 21966
	SSS := time.Duration(21966 * time.Second)
	Future := current.Add(SSS)
	fmt.Println("6 hours, 6 minutes and 6 seconds from now the time will be: ", Future.Format(time.ANSIC))
}
