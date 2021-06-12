package gomarch

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"sync"

	"github.com/husenap/gomarch/gomarch/camera"
	"github.com/husenap/gomarch/gomarch/vec"
)

type Options struct {
	Palette    color.Palette
	FrameCount int
	DeltaTime  float32
	Viewport   image.Rectangle
	FOV        float32

	SDF        func(v vec.Vec3) float32
	CameraTick func(t float32) (position vec.Vec3, lookat vec.Vec3)
}

func doRayMarch(ro, rd vec.Vec3, o *Options) float32 {
	d := float32(0)

	for i := 0; i < 64; i++ {
		p := vec.Add(ro, vec.Scale(rd, d))
		hit := o.SDF(p)
		d += hit
		if hit < 0.001 || d > 100 {
			break
		}
	}

	return d
}

func Render(o Options, out io.Writer) {
	rect := o.Viewport
	width, height := float32(rect.Max.X-rect.Min.X), float32(rect.Max.Y-rect.Min.Y)
	animation := gif.GIF{LoopCount: 0}
	t := float32(0.0)
	wg := sync.WaitGroup{}
	cam := camera.Camera{FOV: o.FOV}

	for i := 0; i < o.FrameCount; i++ {
		frame := image.NewPaletted(rect, o.Palette)
		cam.Update(o.CameraTick(t))

		renderLine := func(y int) {
			defer wg.Done()
			for x := rect.Min.X; x < rect.Max.X; x++ {
				u := float32(x) / width
				v := float32(y) / height
				u = u*2 - 1
				v = v*2 - 1
				u *= width / height

				ro := cam.GetPosition()
				rd := cam.GetRayDirection(u, v)

				d := doRayMarch(ro, rd, &o)

				if d > 100 {
					frame.SetColorIndex(x, y, 0)
				} else {
					frame.SetColorIndex(x, y, 255)
				}
			}
		}
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			wg.Add(1)
			renderLine(y)
		}
		wg.Wait()

		animation.Delay = append(animation.Delay, int(o.DeltaTime/10.0))
		animation.Image = append(animation.Image, frame)

		t += o.DeltaTime / 1000.0
	}

	err := gif.EncodeAll(out, &animation)
	if err != nil {
		fmt.Println(err)
	}
}
