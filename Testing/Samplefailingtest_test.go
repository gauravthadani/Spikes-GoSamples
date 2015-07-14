package SampleTest
import(
"testing"

)

func Add (num1, num2 int) int{
return num1+num2
}



func TestSample(t *testing.T){
	var expected int =3
	var actual int
	actual=Add(2,6)
	if expected != actual {
		t.Error("expected is ",expected," but actual is ", actual)
	}

}
