package main

import (
	"fmt"
	"os"
	"strconv"

	pop "github.com/michael-wang/gopl/ch2/popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Printf("popcount: %v\n", err)
			continue
		}
		fmt.Printf("%d has %d bit(s) of 1\n", x, pop.PopCount(x))
	}
}

// Chap 2.4.1(p. 37)
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
