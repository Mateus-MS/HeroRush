package engine

type Scene struct {
	Update func(engine *Engine)
}