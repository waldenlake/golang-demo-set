package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

func main() {
	limiter := ratelimit.New(100)

	prev := time.Now()

	for i := 0; i < 10; i++ {
		now := limiter.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
