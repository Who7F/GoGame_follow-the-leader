package input

import "github.com/hajimehoshi/ebiten/v2"

type Direction int

const (
	None Direction = iota
	Up
	Down
	Left
	Right
)

type Input interface {
	Update()
	Direction() Direction
	ActionPressed() bool
}

type KeyboardInput struct {
	dir Direction
}

func (k *KeyboardInput) Update() {
	k.dir = None
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		k.dir = Right
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		k.dir = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		k.dir = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		k.dir = Down
	}
}

func (k *KeyboardInput) Direction() Direction {
	return k.dir
}

func (k *KeyboardInput) ActionPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyZ)
}
