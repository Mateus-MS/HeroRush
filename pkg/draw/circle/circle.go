package draw_circle

import rl "github.com/gen2brain/raylib-go/raylib"

func Draw(opts Options) {
	if opts.Radius == 0 {
		return
	}

	rl.DrawCircle(
		int32(opts.Position.X),
		int32(opts.Position.Y),
		opts.Radius,
		opts.Color,
	)

}
