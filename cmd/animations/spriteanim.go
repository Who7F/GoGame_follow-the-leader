package spriteanim

import (
	"follow-the-leader/cmd/input"
	"follow-the-leader/cmd/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animatio struct {
	Frames     [][]*ebiten.Image
	FrameIndex int
	Timer      float64
	Speed      float64
	Loop       bool
	Dir        input.Direction
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

func (a *Animatio) CurrentFrame(dir ui.Direction) *ebiten.Image {
	if len(a.Frames) == 0 {
		return nil
	}
	return a.Frames[a.FrameIndex][dir]
}

func (a *Animatio) Draw(sceen *ebiten.Image, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	sceen.DrawImage(a.CurrentFrame(ui.Direction(a.Dir)), opts)
}
