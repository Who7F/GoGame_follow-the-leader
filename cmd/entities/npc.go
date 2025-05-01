package entities

import (
	"fmt"
	spriteanim "follow-the-leader/cmd/animations"
	"follow-the-leader/cmd/core"
	"follow-the-leader/cmd/maps"
)

// Npc struct
type Npc struct {
	*Sprite
	FollowsLast bool
	State       core.SpriteState
	thingTimer  int
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
		frames, err := LoadSpriteSheet(data.ImagePath, 16, 16, 4)
		if err != nil {
			return nil, fmt.Errorf("failed to load NPC image: %v", err)
		}

		npc := &Npc{
			Sprite: &Sprite{
				X:     data.X,
				Y:     data.Y,
				Dir:   core.None,
				Speed: 1,
				Anim: &spriteanim.Animatio{
					Frames: frames,
					Speed:  0.1,
					Loop:   true,
					Dir:    core.Down,
				},
			},
			FollowsLast: data.FollowsLast,
			State:       core.SpriteState(core.IdleAnim),
			thingTimer:  300,
		}
		npcs = append(npcs, npc)
	}
	return npcs, nil
}

// Update handles NPC movement
func (n *Npc) Update(playerX, playerY float64, colliders []maps.ColliderProvider) {
	i := 20
	if n.thingTimer > i {
		n.thingTimer = 0
		n.Dir = n.Follows(playerX, playerY, 10)
	}
	n.thingTimer++

	if n.Dir != core.None {
		n.Anim.Dir = n.Dir
		n.State = core.SpriteState(core.Walking)
	} else {
		n.State = core.SpriteState(core.IdleAnim)
	}
	n.Anim.State = n.State

	n.Dx = 0
	n.Dy = 0

	if n.FollowsLast {
		switch n.Dir {
		case core.Right:
			n.Dx = n.Speed
		case core.Left:
			n.Dx = -n.Speed
		case core.Up:
			n.Dy = -n.Speed
		case core.Down:
			n.Dy = n.Speed
		}
	} else {
		if n.X > playerX-5 && n.X < playerX+5 && n.Y > playerY-5 && n.Y < playerY+5 {
			n.FollowsLast = true
		}
	}
	n.X += n.Dx

	//CheckCollisionHorizotaly(n.Sprite, colliders)

	n.Y += n.Dy

	//CheckCollisionVertical(n.Sprite, colliders)
}

func (n *Npc) Follows(playerX, playerY, distanceMin float64) core.Direction {

	x := n.X - playerX
	y := n.Y - playerY
	dir := core.None

	if abs(x) > abs(y) {
		if x < -distanceMin {
			dir = core.Right
		} else if x > distanceMin {
			dir = core.Left
		}
	} else {
		if y > distanceMin {
			dir = core.Up
		} else if y < -distanceMin {
			dir = core.Down
		}
	}
	return dir
}

func abs(n float64) float64 {
	if n > 0 {
		return n
	}
	return -n
}
