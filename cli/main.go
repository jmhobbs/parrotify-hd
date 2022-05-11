package main

import (
	"flag"
	"image"
	"image/png"
	"os"

	"github.com/jmhobbs/parrotify-hd/pkg/parrot"
)



func main () {
	var (
		shiftX *int = flag.Int("x", 0, "shift by pixels on the X axis")
		shiftY *int = flag.Int("y", 0, "shift by pixels on the Y axis")
		scale *int = flag.Int("scale", 0, "adjust scale proportionally by pixels in width")
		flip *bool = flag.Bool("flip", false, "flip the overlay image horizontally")
	)
	flag.Parse()

	overlay, err := loadOverlay(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	out, err := parrot.Overlay(overlay, *scale, *shiftX, *shiftY, *flip)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("out.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = os.WriteFile("out.gif", out, 0644); err != nil {
		panic(err)
	}
}

func loadOverlay(path string) (image.Image, error) {
	overlayFile, err := os.Open(flag.Arg(0))
	if err != nil { panic(err) }
	defer overlayFile.Close()

	// todo: input format variations

	return png.Decode(overlayFile)
}