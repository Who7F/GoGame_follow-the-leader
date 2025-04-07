package entities

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Player struct
type Player struct {
	*Sprite
	IsDrunk uint
}

// NewPlayer loads the player sprite
func NewPlayer(x, y float64) (*Player, error) {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		return nil, err
	}

	return &Player{
		Sprite:  &Sprite{Img: img, X: x, Y: y},
		IsDrunk: 0,
	}, nil
}

// Update handles movement
func (p *Player) Update(colliders []image.Rectangle) {
	const playerSpeed float64 = 1
	p.Dx = 0
	p.Dy = 0

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Dx = playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Dx = -playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.Dy = -playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.Dy = playerSpeed
	}

	p.X += p.Dx

	CheckCollisionHorizotaly(p.Sprite, colliders)

	p.Y += p.Dy

	CheckCollisionVertical(p.Sprite, colliders)
}
