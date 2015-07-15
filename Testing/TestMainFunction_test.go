package SampleTest
import (
"testing"
"os"
)


func Add (num1, num2 int) int{
return num1+num2
}

func setup(){

}

func teardown(){

}

func TestMain(m *testing.M) {
	setup()
	retcode := m.Run()
	teardown()
	os.Exit(retcode)
}

func TestSample(t *testing.T){
	var expected int =8
	var actual int
	actual=Add(2,6)
	if expected != actual {
		t.Error("expected is ",expected," but actual is ", actual)
	}
}
