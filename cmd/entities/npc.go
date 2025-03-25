package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Npc struct
type Npc struct {
	*Sprite
	FollowsLast bool
}

// LoadImage
func LoadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load image: %v", err)
	}
	return img, nil
}

func NewNPCs(jsonFile string) ([]*Npc, error) {
	npcData, err := LoadNPCs(jsonFile)
	if err != nil {
		return nil, err
	}
	
	npcs := []*Npc{}
	
	for _, data := range npcData {
		img, err := LoadImage(data.ImagePath)
		if err != nil{
			return nil, fmt.Errorf("failed to load NPC image: %v", err)
		}
		
		npc := &Npc{
			Sprite: &Sprite{
				Img: img,
				X: data.X,
				Y: data.Y,
			},
			FollowsLast: data.FollowsLast,
		}
		npcs = append(npcs, npc)
	}
	return npcs, nil
}

// Update handles NPC movement
func (n *Npc) Update(playerX, playerY float64) {
	if n.FollowsLast {
		if n.X < playerX-5 {
			n.X += 0.5
		} else if n.X > playerX+5 {
			n.X -= 0.5
		}
		if n.Y < playerY-5 {
			n.Y += 0.5
		} else if n.Y > playerY+5 {
			n.Y -= 0.5
		} 
	}else {
		if n.X > playerX - 5 && n.X < playerX + 5 && n.Y > playerY - 5 && n.Y < playerY + 5 {
			n.FollowsLast = true
		}
	}
}
