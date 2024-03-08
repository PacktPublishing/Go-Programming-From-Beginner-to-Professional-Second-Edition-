package main

import (
	"testing"
	"time"
)

func TestElapsed(t *testing.T) {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	result := elapsedTime(start, end)
	t.Log(result)

}
