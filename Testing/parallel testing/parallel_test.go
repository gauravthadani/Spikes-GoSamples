package parallel

import (
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	t.Log("this is test one")
}

func TestTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	t.Log("this is test two")
}


