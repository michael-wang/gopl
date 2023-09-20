package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gerow/go-color"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var (
	width, height float64 = 600, 320
	xyscale               = width / 2 / xyrange
	zscale                = height * 0.4
	sin30, cos30          = math.Sin(angle), math.Cos(angle)
)

func setDim(w, h float64) {
	if w > 0 {
		width = w
	}
	if h > 0 {
		height = h
	}
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("r.Form: %v\n", r.Form)
		setDim(getFloat64(r.Form, "width"), getFloat64(r.Form, "height"))
		setColor(getColor(r.Form, "minZColor"), getColor(r.Form, "maxZColor"))
		w.Header().Set("Content-Type", "image/svg+xml")
		drawSurface(w)
	})
	fmt.Println("server listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getFloat64(form url.Values, key string) float64 {
	v, ok := form[key]
	if !ok {
		return 0
	}
	f, err := strconv.ParseFloat(v[0], 64)
	if err != nil {
		fmt.Printf("parse %s err: %v\n", v, err)
		return 0
	}
	return f
}

func getColor(form url.Values, key string) *color.RGB {
	v := form.Get(key)
	if v == "" {
		return nil
	}
	// v: #FF00FF
	c, err := color.HTMLToRGB(v)
	if err != nil {
		fmt.Printf("HTMLtoRGB err: %v\n", err)
		return nil
	}
	return &c
}

func drawSurface(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", int(width), int(height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			color := mapColor((az + bz + cz + dz) / 4.0)
			fmt.Fprintf(w, "<polygon fill=\"#%s\" points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (sx, sy, z float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z = f(x, y)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

const minZ = -0.3
const maxZ = 1.0

var minZColor = color.RGB{R: 0, G: 0, B: 1.0}.ToHSL() // blue
var maxZColor = color.RGB{R: 1.0, G: 0, B: 0}.ToHSL() // red
// H in HSL color space
var minH = minZColor.H
var maxH = maxZColor.H

func setColor(min, max *color.RGB) {
	fmt.Printf("setColor min: %v, max: %v\n", min, max)
	if min != nil {
		minZColor = min.ToHSL()
	}
	if max != nil {
		maxZColor = max.ToHSL()
	}
	minH = minZColor.H
	maxH = maxZColor.H
	fmt.Printf("in HSL min: %v, max: %v\n", minZColor, maxZColor)
}

// z:		[minZ,		maxZ]
// maxZ-z:	[maxZ-minZ,	0]
// h:		[maxH,		minH]
func mapColor(z float64) string {
	var h float64
	if maxH > minH {
		h = (z-minZ)/(maxZ-minZ)*(maxH-minH) + minH
	} else {
		h = (maxZ-z)/(maxZ-minZ)*(minH-maxH) + maxH
	}
	hsl := color.HSL{H: h, S: 1.0, L: 0.5}
	return hsl.ToHTML()
}
