package spriteanim

import (
	"follow-the-leader/cmd/camera"
	"follow-the-leader/cmd/core"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animatio struct {
	Frames     [][]*ebiten.Image
	FrameIndex int
	Timer      float64
	Speed      float64
	Loop       bool
	Dir        core.Direction
	State      core.SpriteState
}

func (a *Animatio) Update(dt float64) {
	if len(a.Frames) <= 1 {
		return
	} //Only one frame. Dont animate

	a.Timer += dt
	if a.Timer >= a.Speed {
		a.Timer -= a.Speed
		a.FrameIndex++
		if a.FrameIndex >= len(a.Frames[a.Dir]) {
			if a.Loop {
				a.FrameIndex = 0
			} else {
				a.FrameIndex = len(a.Frames[a.Dir]) - 1
			}
		}
	}
}

func (a *Animatio) CurrentFrame(dir core.Direction) *ebiten.Image {
	if len(a.Frames) == 0 {
		return nil
	}
	const tempFix = 1
	// temp fix and Directiion has 5 filds, and the sprietsheet is 4
	const idle = 4
	if core.Direction(a.State) == core.IdleAnim {
		return a.Frames[idle][dir-tempFix]
	}

	return a.Frames[a.FrameIndex][dir-tempFix]
}

func (a *Animatio) Draw(sceen *ebiten.Image, camcam *camera.Camera, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x+camcam.X, y+camcam.Y)
	sceen.DrawImage(a.CurrentFrame(a.Dir), opts)
}
