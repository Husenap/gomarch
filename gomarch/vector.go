package gomarch

type Vector struct {
	X, Y, Z float32
}

func (lhs *Vector) Add(rhs Vector) {
	lhs.X += rhs.X
	lhs.Y += rhs.Y
	lhs.Z += rhs.Z
}
