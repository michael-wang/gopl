package main

import (
	"fmt"

	tempconv "github.com/michael-wang/gopl/ch2/tempconv"
)

func main() {
	fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
