package entities

import (
	"fmt"
	spriteanim "follow-the-leader/cmd/animations"
	"follow-the-leader/cmd/camera"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img          *ebiten.Image
	X, Y, Dx, Dy float64
	Anim         *spriteanim.Animatio
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

func LoadSpriteSheet(path string, frameWidth, frameHeight, frameCount int) ([]*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	config := spriteanim.SpriteSheetConfig{
		FrameWidth:  frameWidth,
		FrameHeight: frameHeight,
		Columns:     4,
		Rows:        7,
	}

	frames := spriteanim.SliceSheet(img, config)
	return frames[0], nil
}

func CheckCollisionHorizotaly(sprite *Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(image.Rect(int(sprite.X), int(sprite.Y), int(sprite.X)+16, int(sprite.Y)+16)) {
			if sprite.Dx > 0.0 {
				sprite.X = float64(collider.Min.X) - 16
			} else if sprite.Dx < 0.0 {
				sprite.X = float64(collider.Max.X)
			}
		}
	}
}

func CheckCollisionVertical(sprite *Sprite, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(image.Rect(int(sprite.X), int(sprite.Y), int(sprite.X)+16, int(sprite.Y)+16)) {
			if sprite.Dy > 0.0 {
				sprite.Y = float64(collider.Min.Y) - 16
			} else if sprite.Dy < 0.0 {
				sprite.Y = float64(collider.Max.Y)
			}
		}
	}
}
