package draw_text

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw(opts Options) {
	if opts.Text == "" {
		return
	}

	if opts.FontSize == 0 {
		opts.FontSize = DefaultTextSize
	}

	switch opts.Aligment {
	case AligmentCenter:
		textWidth := rl.MeasureText(opts.Text, opts.FontSize)
		opts.Position.X -= textWidth / 2
	case AligmentEnd:
		textWidth := rl.MeasureText(opts.Text, opts.FontSize)
		opts.Position.X -= textWidth
	}

	rl.DrawText(
		opts.Text,
		opts.Position.X,
		opts.Position.Y,
		opts.FontSize,
		opts.Color,
	)

}
