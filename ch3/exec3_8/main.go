// Consider using: https://github.com/ALTree/bigfloat for big float/rat calculations like pow or abs.
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	x0, y0        = -2, -2
	x1, y1        = +2, +2
	width, height = 1024, 1024
)

const (
	precC64 = iota
	precC128
	precBigFloat
	precBigRat
)

var precs = map[string]int{
	"c64":      precC64,
	"c128":     precC128,
	"bigfloat": precBigFloat,
	"bigrat":   precBigRat,
}

func main() {
	var (
		x, y float64
		zoom float64
		prec string
		out  string
	)
	flag.Float64Var(&x, "x", 0.0, "x translation")
	flag.Float64Var(&y, "y", 0.0, "y translation")
	flag.Float64Var(&zoom, "z", 1, "zoom in level, e.g. 1, 2, 4...")
	flag.StringVar(&prec, "p", "c64", "precision: c64: complex64, c128: complex128, bigfloat: big.Float, bigrat: big.Rat")
	flag.StringVar(&out, "o", "./out", "output image path")
	flag.Parse()

	file, err := os.Create(out)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}

	draw(file, x, y, zoom, precs[prec])
}

func draw(out io.Writer, x, y, zoom float64, prec int) {
	var (
		xmin = x0/zoom + x
		ymin = y0/zoom + y
		xmax = x1/zoom + x
		ymax = y1/zoom + y
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			switch prec {
			case precC64:
				img.Set(px, py, mandelbrot64(z))
			case precC128:
				img.Set(px, py, mandelbrot128(z))
			case precBigFloat:
				img.Set(px, py, mandelbrotBFloat(z))
			case precBigRat:
				log.Fatal("Unimplemented")
			default:
				log.Fatalf("Invalid precision: %v (valid values: %v)", prec, precs)
			}
		}
	}
	png.Encode(out, img)
}

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		abs := cmplx.Abs(v)
		if abs > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotBFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	zR := new(big.Float).SetFloat64(real(z))
	zI := new(big.Float).SetFloat64(imag(z))
	vR := new(big.Float)
	vI := new(big.Float)
	for n := uint8(0); n < iterations; n++ {
		// v = v*v + z
		vR2, vI2 := new(big.Float), new(big.Float)
		// vR = (vR * vR - vI * vI) + zR
		vR2.Mul(vR, vR).Sub(vR2, new(big.Float).Mul(vI, vI)).Add(vR2, zR)
		// vI = 2 * vR * vI + zI
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2

		// if (vR*vR + vI*vI) > 4 {
		v2 := new(big.Float)
		v2.Mul(vR, vR).Add(v2, new(big.Float).Mul(vI, vI))

		if v2.Cmp(big.NewFloat(4)) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
