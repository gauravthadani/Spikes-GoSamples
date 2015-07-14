package main

import (
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go say("this will never come")
	say("this will come first")
}
