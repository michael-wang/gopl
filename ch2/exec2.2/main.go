package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/michael-wang/gopl/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("exec2.2: %v\n", err)
			os.Exit(1)
		}
		list(v)
	}
}

func list(v float64) {
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
