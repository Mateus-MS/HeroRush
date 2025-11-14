package mathF

import (
	"math"
)

type Vector2 struct {
	X float32
	Y float32
}

func NewVector2(x float32, y float32) Vector2 {
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

func (v Vector2) Mul(scalar float32) Vector2 {
	return Vector2{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vector2) Div(scalar float32) Vector2 {
	return Vector2{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v Vector2) Dot(other Vector2) float32 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2) Len() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2) Normalize() Vector2 {
	len := v.Len()
	if len == 0 {
		return Vector2{0, 0}
	}
	return v.Div(len)
}
