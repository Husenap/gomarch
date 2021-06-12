package vec

import "math"

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func New(x, y, z float32) Vec3 {
	return Vec3{x, y, z}
}
func FromScalar(x float32) Vec3 {
	return Vec3{x, x, x}
}
func Unit() Vec3 {
	return Vec3{1, 1, 1}
}

func Addn(vs ...Vec3) Vec3 {
	var out Vec3
	for _, v := range vs {
		out.X += v.X
		out.Y += v.Y
		out.Z += v.Z
	}
	return out
}
func Add(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs.X + rhs.X, lhs.Y + rhs.Y, lhs.Z + rhs.Z}
}
func Sub(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs.X - rhs.X, lhs.Y - rhs.Y, lhs.Z - rhs.Z}
}
func Scale(v Vec3, f float32) Vec3 {
	return Vec3{v.X * f, v.Y * f, v.Z * f}
}

func Dot(lhs, rhs Vec3) float32 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}
func Length(v Vec3) float32 {
	return float32(math.Sqrt(float64(Dot(v, v))))
}
func Normalize(v Vec3) Vec3 {
	return Scale(v, 1.0/Length(v))
}
func Cross(lhs, rhs Vec3) Vec3 {
	var out Vec3
	out.X = lhs.Y*rhs.Z - lhs.Z*rhs.Y
	out.Y = lhs.Z*rhs.X - lhs.X*rhs.Z
	out.Z = lhs.X*rhs.Y - lhs.Y*rhs.X
	return out
}
