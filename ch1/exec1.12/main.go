package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "err: %v\n", err)
			return
		}
		lissajous(w,
			get(r.Form, "cycles"),
			get(r.Form, "size"),
			get(r.Form, "nframes"),
			get(r.Form, "delay"),
		)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func get(form url.Values, key string) int {
	v, ok := form[key]
	if !ok {
		return 0
	}
	i, err := strconv.Atoi(v[0])
	if err != nil {
		fmt.Printf("parse %s err: %v\n", v, err)
		return 0
	}
	return i
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, cycles, size, nframes, delay int) {
	const res = 0.001

	if cycles == 0 {
		cycles = 5
	}
	if size == 0 {
		size = 100
	}
	if nframes == 0 {
		nframes = 64
	}
	if delay == 0 {
		delay = 8
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
