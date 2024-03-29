package main

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%g lb", p) }

func KToP(k Kilogram) Pound {
	return Pound(k * 2.20462262185)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}
