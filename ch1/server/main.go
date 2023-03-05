package main

import (
	"flag"
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
	"sync"
)

var mu sync.Mutex
var count int
var print bool

func main() {
	flag.BoolVar(&print, "print", false, "server 3: print headers and form data")
	flag.Parse()

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", drawLissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// server 2
	mu.Lock()
	count++
	mu.Unlock()

	// server 1
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// server 3
	if print {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func drawLissajous(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "err: %v\n", err)
		return
	}
	lissajous(w,
		formIntVal(r.Form, "cycles"),
		formIntVal(r.Form, "size"),
		formIntVal(r.Form, "nframes"),
		formIntVal(r.Form, "delay"),
	)
}

// Copy from ../lissajous/main.go
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, cycles, size, nframes, delay int) {
	if cycles == 0 {
		cycles = 5
	}
	const res = 0.001
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
			img.SetColorIndex(
				size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				blackIndex,
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func formIntVal(form url.Values, key string) int {
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
