package camera

import "github.com/husenap/gomarch/gomarch/vec"

type Camera struct {
	FOV      float64
	position vec.Vec3
	right    vec.Vec3
	up       vec.Vec3
	forward  vec.Vec3
}

func (c *Camera) Update(position, lookat vec.Vec3) {
	c.position = position
	c.forward = vec.Normalize(vec.Sub(lookat, position))
	c.right = vec.Normalize(vec.New(c.forward.Z, 0, -c.forward.X))
	c.up = vec.Normalize(vec.Cross(c.forward, c.right))
}

func (c *Camera) GetPosition() vec.Vec3 {
	return c.position
}
func (c *Camera) GetRayDirection(u, v float64) vec.Vec3 {
	return vec.Normalize(
		vec.Addn(
			c.forward,
			vec.Scale(c.right, c.FOV*u),
			vec.Scale(c.up, c.FOV*v)))
}
