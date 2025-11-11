package mathF

import "math"

type Vector2 struct {
	X int32
	Y int32
}

func NewVector2(x int32, y int32) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector2) Mul(scalar int32) Vector2 {
	return Vector2{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vector2) Div(scalar int32) Vector2 {
	return Vector2{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v Vector2) Dot(other Vector2) int32 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2) Len() int32 {
	return int32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2) Normalize() Vector2 {
	len := v.Len()
	if len == 0 {
		return Vector2{0, 0}
	}
	return v.Div(len)
}
