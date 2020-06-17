package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		dwidth, dheight        = width * 2, height * 2
	)

	var superSample [dwidth][dheight]color.Color

	for py := 0; py < dheight; py++ {
		y := float64(py)/dheight*(ymax-ymin) + ymin
		for px := 0; px < dwidth; px++ {
			x := float64(px)/dwidth*(xmax-xmin) + xmin
			z := complex(x, y)
			superSample[px][py] = mandelbrot(z)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			si, sj := 2*i, 2*j

			r1, g1, b1, a1 := superSample[si][sj].RGBA()
			r2, g2, b2, a2 := superSample[si+1][sj].RGBA()
			r3, g3, b3, a3 := superSample[si+1][sj+1].RGBA()
			r4, g4, b4, a4 := superSample[si][sj+1].RGBA()

			// https://golang.org/pkg/image/color/#Color
			// 1024 = 2^8 * 4
			avgColor := color.RGBA{
				uint8((r1 + r2 + r3 + r4) / 1024),
				uint8((g1 + g2 + g3 + g4) / 1024),
				uint8((b1 + b2 + b3 + b4) / 1024),
				uint8((a1 + a2 + a3 + a4) / 1024)}

			img.Set(i, j, avgColor)
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
