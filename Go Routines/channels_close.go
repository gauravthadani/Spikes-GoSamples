package main

import (
"fmt"
)

func factorial(n int, c chan int) {
	y := 1
	for i := n; i > 0; i-- {
		c <- i
		y = i * y
	}
	c <- y
	close(c)
}

func main() {
	c := make(chan int, 3)
	factorial(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}				
}
