package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	// canvas size in pixels
	cells = 100
	// number of grid cells
	xyrange = 30.0
	// axis ranges (-xyrange..+xyrange)
	xyscale = width / 2 / xyrange // pixels per x or y unit
	zscale  = height * 0.4
	// pixels per z unit
	angle = math.Pi / 6
	// angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, b1, z1 := corner(i+1, j)
			bx, by, b2, z2 := corner(i, j)
			cx, cy, b3, z3 := corner(i, j+1)
			dx, dy, b4, z4 := corner(i+1, j+1)
			zAverage := (z1 + z2 + z3 + z4) / 4
			color := lerpRedBlue(zAverage)
			if b1 && b2 && b3 && b4 {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}
func corner(i, j int) (float64, float64, bool, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	if !math.IsNaN(z) {
		sx := width/2 + (x-y)*cos30*xyscale
		sy := height/2 + (x+y)*sin30*xyscale - z*zscale
		return sx, sy, true, z
	}
	return 0, 0, false, 0
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
}
func f(x, y float64) float64 {
	//r := math.Hypot(x, y) // distance from (0,0)
	return 0.06 * (math.Sin(x/2) + math.Sin(y/2))
}
func lerpRedBlue(t float64) string {
	// clamp t to [0,1] to avoid surprises
	t = math.Max(0, math.Min(1, t))

	// Red fades 255→0, Blue grows 0→255, Green stays 0
	r := uint8(math.Round((1 - t) * 255))
	g := uint8(0)
	b := uint8(math.Round(t * 255))

	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
