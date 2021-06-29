package main

import (
	"fmt"
	"math"

	"github.com/gerow/go-color"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			color := mapColor((az + bz + cz + dz) / 4.0)
			fmt.Printf("<polygon fill=\"#%s\" points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

const minZ = -0.3
const maxZ = 1.0

var minZColor = color.RGB{0, 0, 1.0}.ToHSL() // blue
var maxZColor = color.RGB{1.0, 0, 0}.ToHSL() // red
// H in HSL color space
var minH = minZColor.H
var maxH = maxZColor.H

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
	hsl := color.HSL{h, 1.0, 0.5}
	return hsl.ToHTML()
}
