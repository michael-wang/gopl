package main

import (
	"flag"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	backgroundIndex = 0
	foregroundIndex = 1
)

// For exercise 1.5
var palette2 = []color.Color{color.Black, green}
var usePalette2 = flag.Bool("green", false, "Exercise 1.5: change palette to green over black")

// For exercise 1.6
var red = color.RGBA{0xff, 0, 0, 0xff}
var green = color.RGBA{0, 0xff, 0, 0xff}
var blue = color.RGBA{0, 0, 0xff, 0xff}
var palette3 = color.Palette{color.Black, red, green, blue}
var randomColor = flag.Bool("rand", false, "Exercise 1.6: add more colors and SetColorIndex in some ionteresting way")

const sizeOfPalette3 = 4

func main() {
	flag.Parse()

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		plt := palette
		if *usePalette2 {
			plt = palette2
		}
		if *randomColor {
			plt = palette3
		}
		img := image.NewPaletted(rect, plt)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			color := foregroundIndex
			if *randomColor {
				color = rand.Intn(sizeOfPalette3)
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(color))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
