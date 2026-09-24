package engine

type MovementAction struct {
	dx int
	dy int
}

func NewMovementAction(dx int, dy int) MovementAction {
	return MovementAction{dx: dx, dy: dy}
}

func (m MovementAction) Perform(e GameEngine) {
	e.MoveEntity(m.dx, m.dy)
}
