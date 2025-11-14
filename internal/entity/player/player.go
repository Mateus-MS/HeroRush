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
	Position  mathF.Vector2
	Velocity  mathF.Vector2
	WalkSpeed int32
	RunSpeed  int32
	Size      float32

	DashCooldown int
}

func New() Player {
	return Player{
		Position:  mathF.NewVector2(100, 100),
		Velocity:  mathF.NewVector2(0, 0),
		WalkSpeed: 5,
		RunSpeed:  10,
		Size:      15,

		DashCooldown: 100,
	}
}

// TODO: re-write this to preserve the knockback
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

	var targetSpeed float32 = float32(p.WalkSpeed)
	if e.Keys.IsDown("Shift") {
		targetSpeed = float32(p.RunSpeed)
	}

	if dir.X != 0 || dir.Y != 0 {
		dir = dir.Normalize()
		p.Velocity = p.Velocity.Add(dir.Mul(targetSpeed))
	} else {
		p.Velocity = p.Velocity.Mul(0.8)
	}

	// This is a problem with knockbacks
	maxSpeed := float32(p.RunSpeed)
	if p.Velocity.Len() > maxSpeed {
		p.Velocity = p.Velocity.Normalize().Mul(maxSpeed)
	}

	p.Position = p.Position.Add(p.Velocity)
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
				Text:     fmt.Sprintf("Position: %.2f %.2f", p.Position.X, p.Position.Y),
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
				Text:     fmt.Sprintf("Speed: %.2f, %.2f", p.Velocity.X, p.Velocity.Y),
				Position: pos,
				Color:    rl.Black,
				Aligment: draw_text.AligmentStart,
			},
		)
	}
}
