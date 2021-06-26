package main

import (
	"fmt"
	"os"
	"strconv"

	pop "gopl.io/ch2/popcount"
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
