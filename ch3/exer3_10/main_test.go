package main

import "testing"

func TestComma(t *testing.T) {
	s := "1"
	exp := "1"
	testComma(t, s, exp)

	s = "12"
	exp = "12"
	testComma(t, s, exp)

	s = "123"
	exp = "123"
	testComma(t, s, exp)

	s = "1234"
	exp = "1,234"
	testComma(t, s, exp)

	s = "12345"
	exp = "12,345"
	testComma(t, s, exp)

	s = "123456"
	exp = "123,456"
	testComma(t, s, exp)

	s = "1234567"
	exp = "1,234,567"
	testComma(t, s, exp)

	s = "12345678"
	exp = "12,345,678"
	testComma(t, s, exp)

	s = "123456789"
	exp = "123,456,789"
	testComma(t, s, exp)

	s = "1234567890"
	exp = "1,234,567,890"
	testComma(t, s, exp)
}

func testComma(t *testing.T, s, exp string) {
	got := comma(s)
	if exp != got {
		t.Errorf("expected: %s, got: %s\n", exp, got)
	}
}
