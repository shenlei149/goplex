package main

import (
	"os"
	"surface"
)

func main() {
	surface.GenerateSvg(os.Stdout, f)
}

func f(x, y float64) float64 {
	return (x*x+y*y)/surface.XYrange/surface.XYrange*4 - 1
}
