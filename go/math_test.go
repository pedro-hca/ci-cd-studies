package main

import "testing"

func TestSum(t *testing.T) {

	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Sum result invalid: result %d. value expected: %d", total, 30)
	}
}