// Consider using: https://github.com/ALTree/bigfloat for big float/rat calculations like pow or abs.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	x0, y0        = -2, -2
	x1, y1        = +2, +2
	width, height = 1024, 1024
)

func main() {
	var (
		x, y float64 = 0, 0
		zoom float64 = 1
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for k, v := range r.Form {
			var err error
			switch k {
			case "x":
				x, err = strconv.ParseFloat(v[0], 64)
			case "y":
				y, err = strconv.ParseFloat(v[0], 64)
			case "z":
				zoom, err = strconv.ParseFloat(v[0], 64)
			default:
				fmt.Printf("Unknown query param: %s\n", k)
			}

			if err != nil {
				fmt.Println(err)
			}
		}
		draw(w, x, y, zoom)
	})
	fmt.Println("server listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func draw(out io.Writer, x, y, zoom float64) {
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
			img.Set(px, py, mandelbrot128(z))
		}
	}
	png.Encode(out, img)
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
