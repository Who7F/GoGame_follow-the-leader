package entities

import (
	spriteanim "follow-the-leader/cmd/animations"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Player struct
type Player struct {
	*Sprite
	IsDrunk uint
}

// NewPlayer loads the player sprite
func NewPlayer(x, y float64) (*Player, error) {
	//img, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	frames, err := LoadSpriteSheet("assets/images/player.png", 16, 16, 4)
	if err != nil {
		return nil, err
	}

	return &Player{
		Sprite: &Sprite{
			X: x,
			Y: y,
			Anim: &spriteanim.Animatio{
				Frames: frames,
				Speed:  0.1,
				Loop:   true,
			},
		},
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
