package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)

	fmt.Fprintf(os.Stderr, "rBound: %v\n", rBound)
	fmt.Fprintf(os.Stderr, "iBound: %v\n", iBound)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	//const maxRGB = 0xffffff
	//const minRGB = 0x100000

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//c := (maxRGB-minRGB)*uint32(n)/iterations + minRGB
			//rgba := color.RGBA{uint8(c >> 16), uint8(c >> 8), uint8(c), 0xff}
			//fmt.Fprintf(os.Stderr, "c: %X, rgba: %v\n", c, rgba)
			cr := real(v)*0x15 + 0x7e
			cb := imag(v)*0x15 + 0x7e
			c := color.YCbCr{128, uint8(cr), uint8(cb)}
			rBound.check(real(v))
			iBound.check(imag(v))
			return c
		}
	}
	return color.Black
}

type bound struct {
	min float64
	max float64
}

func newBound() *bound {
	return &bound{
		min: math.MaxFloat64,
		max: -math.MaxFloat64,
	}
}

func (b *bound) check(v float64) {
	if v < b.min {
		b.min = v
	}
	if b.max < v {
		b.max = v
	}
}

var rBound = newBound()
var iBound = newBound()
