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

	posX := int32(opts.Position.X)
	posY := int32(opts.Position.Y)

	switch opts.Aligment {
	case AligmentCenter:
		textWidth := rl.MeasureText(opts.Text, opts.FontSize)
		posX -= textWidth / 2
	case AligmentEnd:
		textWidth := rl.MeasureText(opts.Text, opts.FontSize)
		posY -= textWidth
	}

	rl.DrawText(
		opts.Text,
		posX,
		posY,
		opts.FontSize,
		opts.Color,
	)

}
