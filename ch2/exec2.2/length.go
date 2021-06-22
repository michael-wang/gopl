package main

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%g M", m) }
func (f Feet) String() string  { return fmt.Sprintf("%g Ft", f) }

func MToF(m Meter) Feet {
	return Feet(m / 0.3048)
}

func FToM(f Feet) Meter {
	return Meter(f * 0.3048)
}
