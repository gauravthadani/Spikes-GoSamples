package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
			wg.Add(1)
			x := i
			go func() {
				time.Sleep(time.Duration(rand.Int31()))
				fmt.Println(x)
				wg.Done()
			}()
	}
	wg.Wait()
}
