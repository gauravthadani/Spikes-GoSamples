package main
import "fmt"

func sum(a []int, c chan int) {
	sum:=0
	for _,v := range a{
		sum+=v
	}	
	c<-sum
}

func main() {
	a := []int{7,2,8,-9,4,0}

	channel:=make(chan int, 2)
	
	go sum(a[:len(a)/2],channel)
	go sum(a[len(a)/2:],channel)
	x:=<-channel

	go sum(a,channel)
	go sum(a,channel)
	go sum(a,channel)
	
	x=<-channel
	y:=<-channel

	fmt.Println(x)
	fmt.Println(y)
}