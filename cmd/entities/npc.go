package entities

import (
	"fmt"
	"follow-the-leader/cmd/core"
	"image"
)

// Npc struct
type Npc struct {
	*Sprite
	FollowsLast bool
	Speed       float64
	State       core.SpriteState
	Dir         core.Direction
}

type NPCData struct {
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	FollowsLast bool    `json:"followslast"`
	ImagePath   string  `json:"imagePath"`
}

func LoadNPCs(jsonFile string) ([]NPCData, error) {
	return LoadJSON[NPCData](jsonFile, "Rum")
}

func NewNPCs(jsonFile string) ([]*Npc, error) {
	npcData, err := LoadNPCs(jsonFile)
	if err != nil {
		return nil, err
	}

	npcs := []*Npc{}

	for _, data := range npcData {
		img, err := LoadImage(data.ImagePath)
		if err != nil {
			return nil, fmt.Errorf("failed to load NPC image: %v", err)
		}

		npc := &Npc{
			Sprite: &Sprite{
				Img: img,
				X:   data.X,
				Y:   data.Y,
			},
			FollowsLast: data.FollowsLast,
			Speed:       1,
			State:       core.SpriteState(core.IdleAnim),
			Dir:         core.Down,
		}
		npcs = append(npcs, npc)
	}
	return npcs, nil
}

// Update handles NPC movement
func (n *Npc) Update(playerX, playerY float64, colliders []image.Rectangle) {
	n.Dx = 0
	n.Dy = 0
	if n.FollowsLast {
		if n.X < playerX-5 {
			n.Dx = n.Speed
		} else if n.X > playerX+5 {
			n.Dx = -n.Speed
		}
		if n.Y < playerY-5 {
			n.Dy = n.Speed
		} else if n.Y > playerY+5 {
			n.Dy = -n.Speed
		}
	} else {
		if n.X > playerX-5 && n.X < playerX+5 && n.Y > playerY-5 && n.Y < playerY+5 {
			n.FollowsLast = true
		}
	}
	n.X += n.Dx

	CheckCollisionHorizotaly(n.Sprite, colliders)

	n.Y += n.Dy

	CheckCollisionVertical(n.Sprite, colliders)
}

func (n *Npc) Follows(playerX, playerY, distanceMin float64) {
	x := n.X - playerX
	y := n.Y - playerY
	n.Dir = core.None

	if abs(x) > abs(y) {
		if x > distanceMin {
			n.Dir = core.Right
		} else if x < -distanceMin {
			n.Dir = core.Left
		}
	} else {
		if y > distanceMin {
			n.Dir = core.Up
		} else if y < -distanceMin {
			n.Dir = core.Down
		}
	}

}

func abs(n float64) float64 {
	if n > 0 {
		return n
	}
	return -n
}
