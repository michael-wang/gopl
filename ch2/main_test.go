package main

import "testing"

func TestGCD(t *testing.T) {
	x, y, exp := 6, 15, 3
	got := gcd(x, y)
	if got != exp {
		t.Errorf("Expect gcd(%d, %d) = %d but got: %d", x, y, exp, got)
	}

	x, y, exp = 3, 5, 1
	got = gcd(x, y)
	if got != exp {
		t.Errorf("Expect gcd(%d, %d) = %d but got: %d", x, y, exp, got)
	}

	x, y, exp = -6, 15, 3
	// x	y	x%y
	// 15	-6	3
	// -6	3	0
	got = gcd(x, y)
	if got != exp {
		t.Errorf("Expect gcd(%d, %d) = %d but got: %d", x, y, exp, got)
	}

	x, y, exp = -6, -15, -3
	// x	y	x%y
	// -15	-6	-3
	// -6	-3	0
	got = gcd(x, y)
	if got != exp {
		t.Errorf("Expect gcd(%d, %d) = %d but got: %d", x, y, exp, got)
	}
}

func TestFib(t *testing.T) {
	n, exp := 0, 0
	// i	x	y
	//		0	1
	got := fib(n)
	if got != exp {
		t.Errorf("Expect fib(%d) = %d but got : %d", n, exp, got)
	}

	n, exp = 1, 1
	// i	x	y
	// 0	1	1
	got = fib(n)
	if got != exp {
		t.Errorf("Expect fib(%d) = %d but got : %d", n, exp, got)
	}

	n, exp = 2, 1
	// i	x	y
	// 1	1	2
	got = fib(n)
	if got != exp {
		t.Errorf("Expect fib(%d) = %d but got : %d", n, exp, got)
	}

	n, exp = 3, 2
	// i	x	y
	// 2	2	3
	got = fib(n)
	if got != exp {
		t.Errorf("Expect fib(%d) = %d but got : %d", n, exp, got)
	}
}
