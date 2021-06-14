## Basic Example

`main.go`
```go
package main

import (
	"image"
	"math"
	"os"

	"github.com/husenap/gomarch/gomarch"
	"github.com/husenap/gomarch/gomarch/vec"
)

func sdf(position vec.Vec3) float64 {
	d := vec.Length(position) - 1.0
	return d
}

func cameraTick(t float64) (position, lookat vec.Vec3) {
	a := t * 2.0 * math.Pi
	position = vec.Scale(vec.New(math.Cos(a), 0.0, math.Sin(a)), 5.0)
	lookat = vec.Zero()
	return
}

func main() {
	options := gomarch.Options{
		FrameCount: 60,
		DeltaTime:  1000.0 / 60.0,
		Viewport:   image.Rect(0, 0, 160, 90),
		FOV:        0.46,

		SDF:        sdf,
		CameraTick: cameraTick,
	}

	gomarch.Render(options, os.Stdout)
}
```

`Terminal`
```bash
$ go run main.go > out.gif
```

