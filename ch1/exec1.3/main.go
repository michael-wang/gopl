package main

import (
	"fmt"
	"os"
	"strings"
)

func slowJoin(ss []string) string {
	var s, sep string
	for i := 0; i < len(ss); i++ {
		s += sep + ss[i]
		sep = " "
	}
	return s
}

func fastJoin(ss []string) string {
	return strings.Join(ss, " ")
}

func main() {
	fmt.Println("slowJoin:", slowJoin(os.Args[1:]))
	fmt.Println("fastJoin:", fastJoin(os.Args[1:]))
}
