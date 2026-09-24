package engine

import "github.com/hajimehoshi/ebiten/v2"

type Action interface {
	Perform(actor Actor, engine *Engine) error
}

type EscapeAction struct {
}

func NewEscapeAction() *EscapeAction {
	return &EscapeAction{}
}

func (ea *EscapeAction) Perform(actor Actor, engine *Engine) error {
	return ebiten.Termination
}

// Interface Guard
var _ Action = (*EscapeAction)(nil)

type MovementAction struct {
	dx int
	dy int
}

func NewMovementAction(dx int, dy int) *MovementAction {
	return &MovementAction{dx: dx, dy: dy}
}

func (m *MovementAction) Perform(actor Actor, engine *Engine) error {
	entity := actor.GetEntity()
	newX := entity.x + m.dx
	newY := entity.y + m.dy

	if engine.IsOnScreen(newX, newY) {
		entity.x = newX
		entity.y = newY
	}
	return nil
}

// Interface Guard
var _ Action = (*MovementAction)(nil)
