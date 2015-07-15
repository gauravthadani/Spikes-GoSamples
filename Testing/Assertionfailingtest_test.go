package SampleTest
import(
"testing"
"github.com/stretchr/testify/assert"
)

func Add (num1, num2 int) int{
return num1+num2
}



func TestSample(t *testing.T){
	var expected int =7
	assert.Equal(t,expected,Add(2,6))
	}



