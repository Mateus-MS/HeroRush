package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	mathF "github.com/Mateus-MS/HeroRush/pkg/math"
)

type Engine struct {
	Viewport mathF.Vector2
	Scene Scene
	Keys *Input
}

func New(scn Scene) Engine {
	return Engine{
		Viewport: mathF.NewVector2(800, 450),
		Scene: scn,
		Keys: NewInput(),
	}
}

func (e Engine) Run(){
	rl.InitWindow(e.Viewport.X, e.Viewport.Y, "HeroRush")
	defer rl.CloseWindow()

	// Configs
	rl.SetTargetFPS(60)

	// Game Loop
	for !rl.WindowShouldClose() {
		e.Keys.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
        e.Scene.Update(&e)
		rl.EndDrawing()
	}
}