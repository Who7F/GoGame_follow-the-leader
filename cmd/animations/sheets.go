package spriteanim

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheetConfig struct {
	FrameWidth  int
	FrameHeight int
	Columns     int
	Rows        int
}

func SliceSheet(img *ebiten.Image, cfg SpriteSheetConfig) [][]*ebiten.Image {
	frames := [][]*ebiten.Image{}

	for row := 0; row < cfg.Rows; row++ {
		line := []*ebiten.Image{}
		for col := 0; col < cfg.Columns; col++ {
			x := col * cfg.FrameWidth
			y := row * cfg.FrameHeight
			sub := img.SubImage(image.Rect(x, y, x+cfg.FrameWidth, y+cfg.FrameHeight)).(*ebiten.Image)
			line = append(line, sub)
		}
		frames = append(frames, line)
	}

	return frames
}
