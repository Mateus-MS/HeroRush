package draw_circle

import rl "github.com/gen2brain/raylib-go/raylib"

func Draw(opts Options) {
	if opts.Radius == 0 {
		return
	}

	rl.DrawCircle(
		opts.Position.X,
		opts.Position.Y,
		opts.Radius,
		opts.Color,
	)

}
