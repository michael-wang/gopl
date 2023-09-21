package main

import (
	"image"
	"image/color"
	"image/png"
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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

// z^^4 - 1 = 0
func newton(z complex128) color.Color {
	const iterations = 200
	const tolerance = 1e-6

	v := z
	for n := uint8(0); n < iterations; n++ {
		vNew := v - (cmplx.Pow(v, 4)-1)/(4*cmplx.Pow(v, 3))
		abs := cmplx.Abs(vNew - v)
		v = vNew
		if abs < tolerance {
			cr := real(v)*0x15 + 0x7e
			cb := imag(v)*0x15 + 0x7e
			c := color.YCbCr{128, uint8(cr), uint8(cb)}
			return c
		}
	}
	return color.Black
}
