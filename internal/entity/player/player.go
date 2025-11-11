package player

import (
	mathF "github.com/Mateus-MS/HeroRush/pkg/math"
	draw "github.com/Mateus-MS/HeroRush/pkg/draw"
	engine "github.com/Mateus-MS/HeroRush/internal/engine"
)

type Player struct {
	Position mathF.Vector2
	Speed    int32
	Size     float32
}

func New() Player {
	return Player{
		Position: mathF.NewVector2(100, 100),
		Speed: 5,
		Size: 15,
	}
}

func (p *Player) Update(e *engine.Engine){
	dir := mathF.NewVector2(0, 0)
	if e.Keys.IsDown("W") { dir.Y -= 1 }
	if e.Keys.IsDown("S") { dir.Y += 1 }
	if e.Keys.IsDown("A") { dir.X -= 1 }
	if e.Keys.IsDown("D") { dir.X += 1 }

	if dir.X != 0 || dir.Y != 0 {
		dir = dir.Normalize()
		p.Position = p.Position.Add(dir.Mul(p.Speed))
	}
}

func (p Player) Draw(){
	draw.Circle(p.Position, p.Size)
}