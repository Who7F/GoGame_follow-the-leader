package entities

import (
	spriteanim "follow-the-leader/cmd/animations"
	"follow-the-leader/cmd/input"
	"image"
)

// Player struct
type Player struct {
	*Sprite
	IsDrunk uint
	Input   input.Input
	Speed   float64
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
				Dir:    input.Down,
			},
		},
		IsDrunk: 0,
		Speed:   2,
		Input:   &input.KeyboardInput{},
	}, nil
}

// Update handles movement
func (p *Player) Update(colliders []image.Rectangle) {
	p.Input.Update()
	dir := p.Input.Direction()
	p.Dx = 0
	p.Dy = 0

	switch dir {
	case input.Right:
		p.Dx = p.Speed
	case input.Left:
		p.Dx = -p.Speed
	case input.Up:
		p.Dy = -p.Speed
	case input.Down:
		p.Dy = p.Speed
	}

	p.X += p.Dx

	CheckCollisionHorizotaly(p.Sprite, colliders)

	p.Y += p.Dy

	CheckCollisionVertical(p.Sprite, colliders)
}
