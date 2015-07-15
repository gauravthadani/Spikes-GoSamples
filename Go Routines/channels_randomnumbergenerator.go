package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	ch := make(chan int32)

	go func() {
		for ;; {
		time.Sleep(time.Duration((rand.Int31() % 7) * 100) * time.Millisecond)
		x := rand.Int31()
		ch<- x
		}
	}()

	for ;; {
		fmt.Println(<-ch)
	}
}