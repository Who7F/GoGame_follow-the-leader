package entities

import (
	"fmt"
	"follow-the-leader/cmd/camera"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

func (s *Sprite) Draw(screen *ebiten.Image, camcam *camera.Camera) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.X+camcam.X, s.Y+camcam.Y)

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
