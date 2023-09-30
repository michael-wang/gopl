package main

import "testing"

func TestBasename(t *testing.T) {
	s := "a/b/c.go"
	exp := "c"
	testBasename(t, s, exp)

	s = "c.d.go"
	exp = "c.d"
	testBasename(t, s, exp)

	s = "abc"
	exp = "abc"
	testBasename(t, s, exp)
}

func testBasename(t *testing.T, s, exp string) {
	got := basename(s)
	if got != exp {
		t.Errorf("s: %s expected: %s but got: %s\n", s, exp, got)
	}
}

func TestBasename2(t *testing.T) {
	s := "a/b/c.go"
	exp := "c"
	testBasename2(t, s, exp)

	s = "c.d.go"
	exp = "c.d"
	testBasename2(t, s, exp)

	s = "abc"
	exp = "abc"
	testBasename2(t, s, exp)
}

func testBasename2(t *testing.T, s, exp string) {
	got := basename2(s)
	if got != exp {
		t.Errorf("s: %s expected: %s but got: %s\n", s, exp, got)
	}
}
