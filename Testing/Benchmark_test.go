package test

import  (
	"testing"
	"fmt"
)

func BenchmarkSample(b *testing.B) {
	b.StartTimer()
	for i:=0;i<10;i++{
		fmt.Println("hello")
	}
	b.StopTimer()
}
