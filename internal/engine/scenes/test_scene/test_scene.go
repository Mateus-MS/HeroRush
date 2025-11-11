package test_scene

import (
	engine "github.com/Mateus-MS/HeroRush/internal/engine"

	player "github.com/Mateus-MS/HeroRush/internal/entity/player"
)

func New() engine.Scene {
	player := player.New()

	return engine.Scene{
		Update: func(engine *engine.Engine) {
			player.Update(engine)
			player.Draw()
			if engine.DebugMode {
				player.DebugDraw()
			}

			// UI
			drawDashCooldown(player.DashCooldown)
		},
	}
}
