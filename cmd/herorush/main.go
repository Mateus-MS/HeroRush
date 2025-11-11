package main

import (
	engine "github.com/Mateus-MS/HeroRush/internal/engine"
	test_scene "github.com/Mateus-MS/HeroRush/internal/engine/scenes/test_scene"
)

func main() {
	engine.New(test_scene.New()).Run()
}