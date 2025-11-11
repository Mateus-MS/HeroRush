package draw_text

import (
	mathF "github.com/Mateus-MS/HeroRush/pkg/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Aligment int

const (
	AligmentStart Aligment = iota
	AligmentCenter
	AligmentEnd
)

const (
	DefaultTextSize = 15
)

type Options struct {
	Text     string
	Color    rl.Color
	Position mathF.Vector2
	FontSize int32
	Aligment Aligment
}
