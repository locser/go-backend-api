package basic

import "testing"

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 3
	)

	actual := AddOne(1)
	if actual != output {
		t.Errorf("AddOne(%d)", input)
	}
}
