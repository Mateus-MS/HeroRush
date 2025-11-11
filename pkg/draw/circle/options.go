package draw_circle

import (
	mathF "github.com/Mateus-MS/HeroRush/pkg/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Options struct {
	Color    rl.Color
	Position mathF.Vector2
	Radius   float32
}
