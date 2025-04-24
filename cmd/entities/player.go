package entities

import (
	"fmt"
	spriteanim "follow-the-leader/cmd/animations"
	"follow-the-leader/cmd/core"
	"follow-the-leader/cmd/input"
	"follow-the-leader/cmd/maps"
)

// Player struct
type Player struct {
	*Sprite
	IsDrunk uint
	Input   input.Input
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
			X:     x,
			Y:     y,
			Dir:   core.None,
			Speed: 2,
			Anim: &spriteanim.Animatio{
				Frames: frames,
				Speed:  0.1,
				Loop:   true,
				Dir:    core.Down,
			},
		},
		IsDrunk: 0,
		Input:   &input.KeyboardInput{},
		State:   core.SpriteState(core.IdleAnim),
	}, nil
}

// Update handles movement
func (p *Player) Update(colliders []*maps.Colliders) {
	// Get user input
	p.Input.Update()

	p.Sprite.Dir = p.Input.Direction()

	if p.Sprite.Dir != core.None {
		p.Sprite.Anim.Dir = p.Sprite.Dir
		p.State = core.SpriteState(core.Walking)
	} else {
		p.State = core.SpriteState(core.IdleAnim)
	} //why am have p.State and p.spite.anmi.state
	p.Sprite.Anim.State = p.State

	p.Dx = 0
	p.Dy = 0

	switch p.Sprite.Dir {
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

	//CheckCollisionHorizotaly(p.Sprite, colliders)

	p.Y += p.Dy

	//CheckCollisionVertical(p.Sprite, colliders)

	if checkCollision(p.X, p.Y, 16, 16, colliders) {
		fmt.Printf("bump")
	}
}
