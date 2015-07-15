package main
import "fmt"

func sum(a []int, c chan int, shouldBlock bool) {
	sum:=0
	for _,v := range a{
		sum+=v
	}	
	c<-sum

	if shouldBlock {
		fmt.Println("Resuming from Blocking Extracting")
	}
}

func main() {
	a := []int{7,2,8,-9,4,0}

	channel:=make(chan int, 2)
	
	go sum(a[:len(a)/2],channel,false)
	go sum(a[len(a)/2:],channel,false)
	go sum(a,channel ,true)
		
	fmt.Println("Priting the firstResult")
	fmt.Println(<-channel)

	fmt.Println("Printing the secondResult")
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}