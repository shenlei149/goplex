package main

import (
	"fractal"
	"image/color"
	"math/cmplx"
	"os"
)

func main() {
	fractal.GeneratePng(os.Stdout, mandelbrot)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - n, 255 - n*4, 255 - n*16, 255}
		}
	}

	return color.Black
}
