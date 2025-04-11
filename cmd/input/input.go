package input

import (
	"follow-the-leader/cmd/core"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input interface {
	Update()
	Direction() core.Direction
	ActionPressed() bool
}

type KeyboardInput struct {
	dir core.Direction
}

func (k *KeyboardInput) Update() {
	k.dir = core.None
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		k.dir = core.Right
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		k.dir = core.Left
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		k.dir = core.Up
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		k.dir = core.Down
	}
}

func (k *KeyboardInput) Direction() core.Direction {
	return k.dir
}

func (k *KeyboardInput) ActionPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyZ)
}
