package main

import (
	"fmt"
	"math/big"
)

const (
	k  = 1000
	KB = k
	MB = KB * k
	GB = MB * k
	TB = GB * k
	PB = TB * k
	EB = PB * k
	ZB = EB * k
	YB = ZB * k
)

func main() {
	fmt.Printf("KB: %v\n", KB)
	fmt.Printf("MB: %v\n", MB)
	fmt.Printf("GB: %v\n", GB)
	fmt.Printf("TB: %v\n", TB)
	fmt.Printf("PB: %v\n", PB)
	fmt.Printf("EB: %v\n", EB)
	eb, k := big.NewInt(EB), big.NewInt(k)
	fmt.Printf("ZB: %v\n", new(big.Int).Mul(eb, k).String())
	fmt.Printf("YB: %v\n", new(big.Int).Mul(new(big.Int).Mul(eb, k), k).String())
}
