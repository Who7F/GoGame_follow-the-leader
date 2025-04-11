package entities

import (
	spriteanim "follow-the-leader/cmd/animations"
	"follow-the-leader/cmd/core"
	"follow-the-leader/cmd/input"
	"image"
)

// Player struct
type Player struct {
	*Sprite
	IsDrunk uint
	Input   input.Input
	Speed   float64
	State   core.SpriteState
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
				Dir:    core.Down,
			},
		},
		IsDrunk: 0,
		Speed:   2,
		Input:   &input.KeyboardInput{},
		State:   core.SpriteState(core.IdleAnim),
	}, nil
}

// Update handles movement
func (p *Player) Update(colliders []image.Rectangle) {
	p.Input.Update()

	dir := p.Input.Direction()

	if dir != core.None {
		p.Sprite.Anim.Dir = dir
	}

	p.Dx = 0
	p.Dy = 0

	switch dir {
	case core.Right:
		p.Dx = p.Speed
	case core.Left:
		p.Dx = -p.Speed
	case core.Up:
		p.Dy = -p.Speed
	case core.Down:
		p.Dy = p.Speed
	}

	p.X += p.Dx

	CheckCollisionHorizotaly(p.Sprite, colliders)

	p.Y += p.Dy

	CheckCollisionVertical(p.Sprite, colliders)
}
