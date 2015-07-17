package parallel

import(
	"testing"
	"time"
)

func TestSampleOne(t *testing.T){
	t.Parallel()
	time.Sleep(2*time.Second)
	t.Log("this is test sample one")
}

func TestSampleTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(2*time.Second)
	t.Log("this is test sample two")
}
//This is not parallel , this will run before the previous two
func TestSampleThree(t *testing.T){
	t.Log("this is test sample three")
}

