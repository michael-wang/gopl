package tempconv

import "fmt"

func ExampleOne() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // 100°C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // 180°F
	/*
		fmt.Printf("%g\n", boilingF-FreezingC)       // compiler error
	*/
}

func ExampleTwo() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	/*
		fmt.Println(c == f)          // compiler error
	*/
	fmt.Println(c == Celsius(f)) // true
}

func MoreExample() {
	c := FToC(212.0)
	fmt.Println(c.String()) // 100°C
	fmt.Printf("%v\n", c)   // 100°C
	fmt.Printf("%s\n", c)   // 100°C
	fmt.Println(c)          // 100°C
	fmt.Printf("%g\n", c)   // 100, does not call String
	fmt.Println(float64(c)) // 100, does not call String
}
