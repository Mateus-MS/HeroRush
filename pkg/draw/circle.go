package draw

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	mathF "github.com/Mateus-MS/HeroRush/pkg/math"
)

func Circle(position mathF.Vector2, radius float32){
	rl.DrawCircle(position.X, position.Y, radius, rl.Red)
}