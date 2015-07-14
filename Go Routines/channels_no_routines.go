package main

import (
"fmt"
)

func main() {
	c := make(chan int, 3)
	
	 func(){
		c <- 3
		c <- 4
		}()

	 func(){
			x:= <-c
			fmt.Println(x)
			close(c)
			}()

			for i := range c {
				fmt.Println(i)
			}	

		}
