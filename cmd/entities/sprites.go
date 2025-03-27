package entities

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.Img.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), opts)
}

// LoadImage
func LoadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load image: %v", err)
	}
	return img, nil
}
