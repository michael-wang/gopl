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
		fmt.Println("from arguments")
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	} else {
		fmt.Println("from stdin")
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			convert(input.Text())
		}
	}
}

func convert(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
