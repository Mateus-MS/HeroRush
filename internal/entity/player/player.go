package player

import (
	"fmt"

	engine "github.com/Mateus-MS/HeroRush/internal/engine"
	draw_circle "github.com/Mateus-MS/HeroRush/pkg/draw/circle"
	draw_text "github.com/Mateus-MS/HeroRush/pkg/draw/text"
	mathF "github.com/Mateus-MS/HeroRush/pkg/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position mathF.Vector2
	Speed    int32
	Size     float32

	DashCooldown int
}

func New() Player {
	return Player{
		Position: mathF.NewVector2(100, 100),
		Speed:    5,
		Size:     15,

		DashCooldown: 100,
	}
}

func (p *Player) Update(e *engine.Engine) {
	dir := mathF.NewVector2(0, 0)
	if e.Keys.IsDown("W") {
		dir.Y -= 1
	}
	if e.Keys.IsDown("S") {
		dir.Y += 1
	}
	if e.Keys.IsDown("A") {
		dir.X -= 1
	}
	if e.Keys.IsDown("D") {
		dir.X += 1
	}

	if dir.X != 0 || dir.Y != 0 {
		dir = dir.Normalize()
		p.Position = p.Position.Add(dir.Mul(p.Speed))
	}

	if e.Keys.IsPressed("Space") {
		println("Apertou o dash")
	}
}

func (p Player) Draw() {
	draw_circle.Draw(
		draw_circle.Options{
			Position: p.Position,
			Color:    rl.Red,
			Radius:   p.Size,
		},
	)
}

func (p Player) DebugDraw() {
	{
		pos := p.Position
		pos.X -= 60
		pos.Y -= 40

		draw_text.Draw(
			draw_text.Options{
				Text:     fmt.Sprintf("Position: %d %d", p.Position.X, p.Position.Y),
				Position: pos,
				Color:    rl.Black,
				Aligment: draw_text.AligmentStart,
			},
		)
	}

	{
		pos := p.Position
		pos.X -= 60
		pos.Y -= 60

		draw_text.Draw(
			draw_text.Options{
				Text:     fmt.Sprintf("Speed: %d", p.Speed),
				Position: pos,
				Color:    rl.Black,
				Aligment: draw_text.AligmentStart,
			},
		)
	}
}
