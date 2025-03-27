package entities

import (
	"fmt"
)

// Rum struct
type Rum struct {
	*Sprite
	GiveDrunk uint
}

// NewRums initializes Rum bottles
func NewRums(jsonFile string) ([]*Rum, error) {
	rumData, err := LoadRums(jsonFile)
	if err != nil {
		return nil, err
	}

	rums := []*Rum{}

	for _, data := range rumData {
		img, err := LoadImage(data.ImagePath)
		if err != nil {
			return nil, fmt.Errorf("failed to load Rum image: %v", err)
		}

		rum := &Rum{
			Sprite: &Sprite{
				Img: img,
				X:   data.X,
				Y:   data.Y,
			},
			GiveDrunk: data.GiveDrunk,
		}
		rums = append(rums, rum)
	}

	return rums, nil

}
