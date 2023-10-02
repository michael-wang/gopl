package main

import (
	"testing"
)

func TestPrintints(t *testing.T) {
	values := []int{1, 2, 3}
	exp := "[1, 2, 3]"
	got := intsToString(values)
	t.Logf("got: %s\n", got)
	if exp != got {
		t.Errorf("expected: %s, got: %s\n", exp, got)
	}
}
