package engine

type MovementAction struct {
	dx int
	dy int
}

func NewMovementAction(dx int, dy int) MovementAction {
	return MovementAction{dx: dx, dy: dy}
}

func (m *MovementAction) IsAction() bool {
	return true
}

func (e *Engine) MoveEntity(dx int, dy int) {
	newX := e.playerX + dx
	newY := e.playerY + dy
	if e.IsOnScreen(newX, newY) {
		e.playerX += dx
		e.playerY += dy
	}
}

func (m *MovementAction) Perform(e *Engine) {
	e.MoveEntity(m.dx, m.dy)
}
