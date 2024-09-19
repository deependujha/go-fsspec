package go_fsspec

import "testing"

func TestHello(t *testing.T) {
	a := 1
	b := 1
	c := a + b
	if c != 2 {
		t.Error("something is wrong!")
	}
}
