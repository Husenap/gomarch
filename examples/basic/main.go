package main

import (
	"image"
	"image/color"
	"os"

	"github.com/husenap/gomarch/gomarch"
)

func generatePalette() color.Palette {
	palette := []color.Color{}
	for i := 0; i < 256; i++ {
		palette = append(palette, color.RGBA{uint8(i), uint8(i), 255, 255})
	}
	return palette
}

func sdf(x, y int) int {
	return x ^ y
}

func main() {
	options := gomarch.Options{
		Palette:    generatePalette(),
		FrameCount: 64,
		DeltaTime:  100.0,
		SDF:        sdf,
		Viewport:   image.Rect(0, 0, 128, 128),
	}

	gomarch.Render(options, os.Stdout)
}
