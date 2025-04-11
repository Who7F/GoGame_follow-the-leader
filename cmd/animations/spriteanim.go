package spriteanim

import (
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
	return a.Frames[a.FrameIndex][dir-tempFix]
}

func (a *Animatio) Draw(sceen *ebiten.Image, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	sceen.DrawImage(a.CurrentFrame(a.Dir), opts)
}
