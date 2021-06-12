package vec

import (
	"image/color"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func New(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}
func FromScalar(x float64) Vec3 {
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
func Scale(v Vec3, f float64) Vec3 {
	return Vec3{v.X * f, v.Y * f, v.Z * f}
}

func Dot(lhs, rhs Vec3) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}
func Length(v Vec3) float64 {
	return float64(math.Sqrt(float64(Dot(v, v))))
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

func ToColor(v Vec3) color.RGBA {
	return color.RGBA{
		uint8(math.Max(math.Min(v.X, 1.0), 0.0) * 255),
		uint8(math.Max(math.Min(v.Y, 1.0), 0.0) * 255),
		uint8(math.Max(math.Min(v.Z, 1.0), 0.0) * 255),
		255,
	}
}
