package main

import (
	"math/rand"
	"testing"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
)

func z() complex128 {
	return complex(rand.Float64()*(xmax-xmin)+xmin, rand.Float64()*(ymax-ymin)*ymin)
}

func BenchmarkMandelbrot64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot64(z())
	}
}

func BenchmarkMandelbrot128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot128(z())
	}
}

func BenchmarkMandelbrotBFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotBFloat(z())
	}
}
