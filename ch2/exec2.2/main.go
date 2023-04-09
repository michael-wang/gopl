package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/michael-wang/gopl/ch2/tempconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			list(arg)
			fmt.Println("----------------------")
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			list(input.Text())
			fmt.Println("----------------------")
		}
	}
}

func list(s string) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("exec2.2: %v\n", err)
		os.Exit(1)
	}
	c := tempconv.Celsius(v)
	fmt.Printf("Temparature: %s\n", c)
	fmt.Printf("Temparature: %s\n", tempconv.CToF(c))
	m := Meter(v)
	fmt.Printf("Length: %s\n", m)
	fmt.Printf("Length: %s\n", MToF(m))
	k := Kilogram(v)
	fmt.Printf("Weight: %s\n", k)
	fmt.Printf("Weight: %s\n", KToP(k))
}
