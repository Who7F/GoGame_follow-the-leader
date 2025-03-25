package entities

import (
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
		Sprite: &Sprite{Img: img, X: x, Y: y},
		IsDrunk: 0,
	}, nil
}

// Update handles movement
func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.X += 0.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.X -= 0.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.Y -= 0.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.Y += 0.5
	}
}
