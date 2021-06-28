package main

import (
	"fmt"
	"math"
)

func format() {
	for x := 0; x < 16; x++ {
		fmt.Printf("x = %d\te^x = %11.3f\n", x, math.Exp(float64(x)))
	}
}

func specialValues() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z, math.IsNaN(z/z))
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}

func compute() (result float64, ok bool) {
	//...
	failed := true
	if failed {
		return 0, false
	}
	return result, true
}

func main() {
	format()
	specialValues()
	result, ok := compute()
	fmt.Println("compute result:", result, "ok:", ok)
}
