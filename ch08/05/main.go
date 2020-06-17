package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

type input struct {
	y  float64
	py int
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 4096, 4096
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	numcore := 4
	var ins []chan input
	var wg sync.WaitGroup

	for i := 0; i < numcore; i++ {
		ins = append(ins, make(chan input, height))
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for {
				input, ok := <-ins[index]
				if !ok {
					break
				}
				y := input.y
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					c := mandelbrot(z)
					img.Set(px, input.py, c)
				}
			}
		}(i)
	}

	start := time.Now()
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		ins[py%numcore] <- input{y, py}
	}
	for _, in := range ins {
		close(in)
	}

	wg.Wait()

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	end := time.Now()
	log.Println(end.Sub(start))
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
