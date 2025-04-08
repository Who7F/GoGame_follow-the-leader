package spriteanim

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Animatio struct {
	Frames     []*ebiten.Image
	FrameIndex int
	Timer      float64
	Speed      float64
	Loop       bool
}

func (a *Animatio) Update(dt float64) {
	if len(a.Frames) <= 1 {
		return
	} //Only one frame. Dont animate

	a.Timer += dt
	if a.Timer >= a.Speed {
		a.Timer -= a.Speed
		a.FrameIndex++
		if a.FrameIndex >= len(a.Frames) {
			if a.Loop {
				a.FrameIndex = 0
			} else {
				a.FrameIndex = len(a.Frames) - 1
			}
		}
	}
}

func (a *Animatio) CurrentFrame() *ebiten.Image {
	if len(a.Frames) == 0 {
		return nil
	}
	return a.Frames[a.FrameIndex]
}

func (a *Animatio) Draw(sceen *ebiten.Image, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	sceen.DrawImage(a.CurrentFrame(), opts)
}
