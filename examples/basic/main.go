package main

import (
	"image"
	"math"
	"os"

	"github.com/husenap/gomarch/gomarch"
	"github.com/husenap/gomarch/gomarch/vec"
)

func sdf(position vec.Vec3) float64 {
	return vec.Length(position) - 1.0
}

func cameraTick(t float64) (position, lookat vec.Vec3) {
	x := math.Sin(t*2.0*math.Pi) * 3.0
	position = vec.New(x, 0, -5)
	lookat = vec.New(x, 0, 0)
	return
}

func main() {
	options := gomarch.Options{
		FrameCount: 120,
		DeltaTime:  1000.0 / 60.0,
		Viewport:   image.Rect(0, 0, 160, 90),
		FOV:        0.46,

		SDF:        sdf,
		CameraTick: cameraTick,
	}

	gomarch.Render(options, os.Stdout)
}
