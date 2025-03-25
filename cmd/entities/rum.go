package entities

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Rum struct
type Rum struct {
	*Sprite
	GiveDrunk uint
}

// NewRums initializes Rum bottles
func NewRums() []*Rum {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/rum.png")
	if err != nil {
		panic(err)
	}

	return []*Rum{
		{Sprite: &Sprite{Img: img, X: 100, Y: 100}, GiveDrunk: 10},
	}
}
