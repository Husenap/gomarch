package gomarch

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"io"
	"sync"

	"github.com/husenap/gomarch/gomarch/camera"
	"github.com/husenap/gomarch/gomarch/vec"
)

type Options struct {
	FrameCount int
	DeltaTime  float64
	Viewport   image.Rectangle
	FOV        float64

	SDF        func(v vec.Vec3) float64
	CameraTick func(t float64) (position vec.Vec3, lookat vec.Vec3)
}

type RenderContext struct {
	u   float64
	v   float64
	cam *camera.Camera
	o   *Options
}

func doRayMarch(ro, rd vec.Vec3, o *Options) float64 {
	d := float64(0)

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

const (
	e = 0.0003
)

var e1 = vec.New(e, -e, -e)
var e2 = vec.New(-e, -e, e)
var e3 = vec.New(-e, e, -e)
var e4 = vec.New(e, e, e)

func calcNormal(p vec.Vec3, o *Options) vec.Vec3 {
	return vec.Normalize(
		vec.Addn(
			vec.Scale(e1, o.SDF(vec.Add(p, e1))),
			vec.Scale(e2, o.SDF(vec.Add(p, e2))),
			vec.Scale(e3, o.SDF(vec.Add(p, e3))),
			vec.Scale(e4, o.SDF(vec.Add(p, e4)))))
}

func doRender(rc RenderContext) vec.Vec3 {
	ro := rc.cam.GetPosition()
	rd := rc.cam.GetRayDirection(rc.u, rc.v)

	d := doRayMarch(ro, rd, rc.o)

	if d > 100 {
		return vec.FromScalar(0)
	}

	p := vec.Add(ro, vec.Scale(rd, d))
	n := calcNormal(p, rc.o)

	diffuse := vec.Dot(n, vec.Normalize(vec.New(1, 1, -1)))

	light := diffuse

	return vec.Scale(vec.New(0.4, 0.8, 0.95), light)
}

func Render(o Options, out io.Writer) {
	rect := o.Viewport
	width, height := float64(rect.Max.X-rect.Min.X), float64(rect.Max.Y-rect.Min.Y)
	animation := gif.GIF{LoopCount: 0}
	t := 0.0
	wg := sync.WaitGroup{}
	cam := camera.Camera{FOV: o.FOV}

	for i := 0; i < o.FrameCount; i++ {
		frame := image.NewRGBA(rect)
		palettedFrame := image.NewPaletted(rect, palette.Plan9)

		cam.Update(o.CameraTick(t))

		renderLine := func(y int) {
			defer wg.Done()
			for x := rect.Min.X; x < rect.Max.X; x++ {
				u := float64(x) / width
				v := float64(y) / height
				u = u*2 - 1
				v = -(v*2 - 1)
				u *= width / height

				rc := RenderContext{
					u:   u,
					v:   v,
					cam: &cam,
					o:   &o,
				}
				frame.SetRGBA(x, y, vec.ToColor(doRender(rc)))
			}
		}
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			wg.Add(1)
			go renderLine(y)
		}
		wg.Wait()

		draw.FloydSteinberg.Draw(palettedFrame, rect, frame, image.Point{0, 0})

		animation.Delay = append(animation.Delay, int(o.DeltaTime/10.0))
		animation.Image = append(animation.Image, palettedFrame)

		t += o.DeltaTime / 1000.0
	}

	err := gif.EncodeAll(out, &animation)
	if err != nil {
		fmt.Println(err)
	}
}
