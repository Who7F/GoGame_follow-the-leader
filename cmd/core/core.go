package core

type Direction int

const (
	None Direction = iota
	Down
	Up
	Left
	Right
)

type SpriteState int

const (
	IdleAnim Direction = iota
	Walking
)

type UI struct {
	X, Y int
}
