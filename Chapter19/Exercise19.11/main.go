package main

import (
	"context"
	"log"
	"time"
)

func main() {
	r := make(chan int)
	ctx := context.TODO()
	cl, stop := context.WithCancel(ctx)

	go countNumbers(cl, r)
	go func() {
		time.Sleep(time.Millisecond * 100 * 3)
		stop()
	}()

	v := <-r
	log.Println(v)
}
func countNumbers(ctx context.Context, r chan int) {
	v := 0
	for {
		select {
		case <-ctx.Done():
			r <- v
			break
		default:
			time.Sleep(time.Millisecond * 100)
			v++
		}
	}
}
