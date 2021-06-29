package main

import "fmt"

func complexNumbers() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Printf("%v x %v = %v\n", x, y, x*y)
	fmt.Printf("real of %v x %v = %v\n", x, y, real(x*y))
	fmt.Printf("imag of %v x %v = %v\n", x, y, imag(x*y))

	fmt.Println(1i * 1i)

	fmt.Printf("x := 1 + 2i, x = %v\n", 1+2i)
	fmt.Printf("y := 3.0 + 4i, x = %v\n", 3.0+4i)
}

func main() {
	complexNumbers()
}
