package main

import "testing"

func TestComma(t *testing.T) {
	s := "12345"
	exp := "12,345"
	testComma(t, s, exp)

	s = "123"
	exp = "123"
	testComma(t, s, exp)

	s = "1234567"
	exp = "1,234,567"
	testComma(t, s, exp)

	s = "1"
	exp = "1"
	testComma(t, s, exp)
}

func testComma(t *testing.T, s, exp string) {
	got := comma(s)
	if exp != got {
		t.Errorf("s: %s expect: %s but got: %s\n", s, exp, got)
	}
}
