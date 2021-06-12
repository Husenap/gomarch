package main

import (
	"image/color"
	"os"

	"github.com/husenap/go-raymarching/raymarching"
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
	options := raymarching.Options{
		Palette:    generatePalette(),
		FrameCount: 64,
		DeltaTime:  100.0,
		SDF:        sdf,
	}

	raymarching.Render(options, os.Stdout)
}
