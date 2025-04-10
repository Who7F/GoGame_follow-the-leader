package core

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
