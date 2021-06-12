package main

import (
	"image"
	"image/color"
	"math"
	"os"

	"github.com/husenap/gomarch/gomarch"
	"github.com/husenap/gomarch/gomarch/vec"
)

func generatePalette() color.Palette {
	palette := []color.Color{}
	for i := 0; i < 256; i++ {
		palette = append(palette, color.RGBA{255, uint8(i), uint8(i), 255})
	}
	return palette
}

func sdf(position vec.Vec3) float32 {
	return vec.Length(position) - 1.0
}

func cameraTick(t float32) (position, lookat vec.Vec3) {
	x := float32(math.Sin(float64(t*2.0*math.Pi))) * 3.0
	position = vec.New(x, 0, -5)
	lookat = vec.New(x, 0, 0)
	return
}

func main() {
	options := gomarch.Options{
		Palette:    generatePalette(),
		FrameCount: 120,
		DeltaTime:  1000.0 / 60.0,
		Viewport:   image.Rect(0, 0, 160, 90),
		FOV:        0.46,

		SDF:        sdf,
		CameraTick: cameraTick,
	}

	gomarch.Render(options, os.Stdout)
}
