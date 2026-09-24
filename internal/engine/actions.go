package engine

type GameEngine interface {
	MoveEntity(dx int, dy int)
	IsOnScreen(x int, y int) bool
}

type Action interface {
	Perform(e GameEngine)
}
