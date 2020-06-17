package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func main() {
	var format string
	flag.StringVar(&format, "t", "", "output image type. png, jpg, or gif.")
	flag.Parse()

	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	format = strings.ToUpper(format)
	switch format {
	case "JPG", "JPEG":
		err = jpeg.Encode(os.Stdout, img, nil)
	case "PNG":
		err = png.Encode(os.Stdout, img)
	case "GIF":
		err = gif.Encode(os.Stdout, img, nil)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
