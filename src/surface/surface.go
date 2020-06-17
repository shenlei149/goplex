package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	XYrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / XYrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func GenerateSvg(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ah := corner(i+1, j, f)
			bx, by, bh := corner(i, j, f)
			cx, cy, ch := corner(i, j+1, f)
			dx, dy, dh := corner(i+1, j+1, f)
			var eh float64 = -1.0
			if i > 0 && j > 0 {
				_, _, eh = corner(i-1, j-1, f)
			}

			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}

			var color string
			if bh > ah && bh > ch && bh > dh && bh > eh {
				color = "#ff0000"
			} else if bh < ah && bh < ch && bh < dh && bh < eh {
				color = "#0000ff"
			}

			fmt.Fprintf(w, "<polygon style='stroke:%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (sx, sy, z float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := XYrange * (float64(i)/cells - 0.5)
	y := XYrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z = f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}
