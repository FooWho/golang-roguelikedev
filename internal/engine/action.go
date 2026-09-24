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
	newX := actor.GetX() + m.dx
	newY := actor.GetY() + m.dy

	if engine.IsOnScreen(newX, newY) {
		actor.SetPosition(newX, newY)
	}
	return nil
}

// Interface Guard
var _ Action = (*MovementAction)(nil)

type WaitAction struct {
}

func NewWaitAction() *WaitAction {
	return &WaitAction{}
}

func (w *WaitAction) Perform(actor Actor, engine *Engine) error {
	return nil
}

// Interface Guard
var _ Action = (*WaitAction)(nil)
