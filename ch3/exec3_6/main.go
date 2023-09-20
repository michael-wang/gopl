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

var (
	rBound = newBound()
	iBound = newBound()
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		// y := float64(py)/height*(ymax-ymin) + ymin
		y0 := (float64(py)-0.5)/height*(ymax-ymin) + ymin
		y1 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		y2 := (float64(py)-0.5)/height*(ymax-ymin) + ymin
		y3 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			// x := float64(px)/width*(xmax-xmin) + xmin
			x1 := (float64(px)-0.5)/width*(xmax-xmin) + xmin
			x0 := (float64(px)+0.5)/width*(xmax-xmin) + xmin
			x2 := (float64(px)-0.5)/width*(xmax-xmin) + xmin
			x3 := (float64(px)+0.5)/width*(xmax-xmin) + xmin

			// z := complex(x, y)
			z0 := complex(x0, y0)
			z1 := complex(x1, y1)
			z2 := complex(x2, y2)
			z3 := complex(x3, y3)

			c0 := mandelbrot(z0)
			c1 := mandelbrot(z1)
			c2 := mandelbrot(z2)
			c3 := mandelbrot(z3)

			img.Set(px, py, avgColor([]color.Color{c0, c1, c2, c3}))
		}
	}
	png.Encode(os.Stdout, img)

	fmt.Fprintf(os.Stderr, "rBound: %v\n", rBound)
	fmt.Fprintf(os.Stderr, "iBound: %v\n", iBound)
}

func avgColor(colors []color.Color) color.Color {
	var avgR, avgG, avgB, avgA uint32
	for _, c := range colors {
		r, g, b, a := c.RGBA()
		avgR += r
		avgG += g
		avgB += b
		avgA += a
	}
	n := uint32(len(colors))
	avgR /= n
	avgG /= n
	avgB /= n
	avgA /= n

	return color.RGBA{
		R: uint8(avgR >> 8),
		G: uint8(avgG >> 8),
		B: uint8(avgB >> 8),
		A: uint8(avgA >> 8),
	}
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
