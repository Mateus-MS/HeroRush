package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Input struct {
	Key  map[string]int32
	Keys map[int32]bool
}

func NewInput() *Input {
	return &Input{
		Key: createKeyMap(),
		Keys: make(map[int32]bool),
	}
}

func createKeyMap() map[string]int32 {
	keyMap := make(map[string]int32)

	keyMap["W"] = rl.KeyW
	keyMap["A"] = rl.KeyA
	keyMap["S"] = rl.KeyS
	keyMap["D"] = rl.KeyD

	return keyMap
}

func (i *Input) Update(){
	i.Keys[rl.KeyW] = rl.IsKeyDown(rl.KeyW)
	i.Keys[rl.KeyA] = rl.IsKeyDown(rl.KeyA)
	i.Keys[rl.KeyS] = rl.IsKeyDown(rl.KeyS)
	i.Keys[rl.KeyD] = rl.IsKeyDown(rl.KeyD)
}

func (i *Input) IsDown(action string) bool {
	return i.Keys[i.Key[action]]
}