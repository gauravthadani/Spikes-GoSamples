package main

import(
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go PumpData(ch)
	recievedInt := <- ch
	fmt.Println(recievedInt)
}

func PumpData(ch chan int) {
	time.Sleep(5 * time.Second)	
	ch <- 5	
}