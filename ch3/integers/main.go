package main

import "fmt"

func bitwiseOperators() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
}

func printIntegers() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}

// rune type represent a Unicode code point.
func runes() {
	ascii := 'a'
	unicode := '國'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
}

func sizeType() {
	// it seems len should return uint since length could not
	// less than zero.
	// let's try it, redefine len which returns uint.
	len := func(slice []string) uint {
		return uint(cap(slice))
	}

	// let's try traverse string slice in reverse order.
	medals := []string{"gold", "silver", "bronze"}

	// i will be out of index for 4th iteration.
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Printf("medals[%d]: ", i)
		fmt.Println(medals[i])
	}

	// to fix above issue, need ugly twick on index value.
	for i := len(medals); i > 0; i-- {
		j := i - 1
		fmt.Printf("medals[%d]: ", j)
		fmt.Println(medals[j])
	}

	// summary: that's why len or cap return int instead of uint.
}

func main() {
	bitwiseOperators()
	printIntegers()
	runes()
	sizeType()
}
