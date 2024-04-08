package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	r := rng.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}
