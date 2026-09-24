package main

type Action interface {
	IsAction() bool
}

type MovementAction struct {
	dx int
	dy int
}

type EscapeAction struct {
}

func NewEscapeAction() EscapeAction {
	return EscapeAction{}
}

func (e *EscapeAction) IsAction() bool {
	return true
}

func NewMovementAction(dx int, dy int) MovementAction {
	ma := MovementAction{dx: dx, dy: dy}
	return ma
}

func (m *MovementAction) IsAction() bool {
	return true
}
