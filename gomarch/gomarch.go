package gomarch

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
)

type Options struct {
	Palette    color.Palette
	FrameCount int
	DeltaTime  float32
	SDF        func(x, y int) int
	Viewport   image.Rectangle
}

func Render(o Options, out io.Writer) {
	rect := o.Viewport
	animation := gif.GIF{LoopCount: 0}

	for i := 0; i < o.FrameCount; i++ {
		frame := image.NewPaletted(rect, o.Palette)

		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			for x := rect.Min.X; x < rect.Max.X; x++ {
				frame.SetColorIndex(x, y, uint8(o.SDF(x, y)))
			}
		}

		animation.Delay = append(animation.Delay, int(o.DeltaTime/10.0))
		animation.Image = append(animation.Image, frame)
	}

	err := gif.EncodeAll(out, &animation)
	if err != nil {
		fmt.Println(err)
	}
}
