package ui

type Direction int

const (
	Down Direction = iota
	Up
	Left
	Right
)

type UI struct {
	X, Y int
}

//func (u *UI) DrawDirectionArrow(screen *ebiten.Image, direction Direction) {
//	var arrowImage *ebiten.Image
//	switch direction {
//	case Down:
//		arrowImage = u.arrowDown
//	case Up:
//		arrowImage = u.arrowUp
//	case Left:
//		arrowImage = u.arrowLeft
//	case Right:
//		arrowImage = u.arrowRight
//	}
//
//	screen.DrawImage(arrowImage, nil)
//}
