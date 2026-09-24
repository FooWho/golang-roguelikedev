package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (e *Engine) GetAction() Action {
	e.keys = inpututil.AppendJustPressedKeys(e.keys[:0])

	if len(e.keys) == 0 {
		return nil
	}

	switch e.keys[0] {
	case ebiten.KeyArrowUp, ebiten.KeyW:
		movementAction := NewMovementAction(0, -1)
		return &movementAction
	case ebiten.KeyArrowDown, ebiten.KeyS:
		movementAction := NewMovementAction(0, 1)
		return &movementAction
	case ebiten.KeyArrowLeft, ebiten.KeyA:
		movementAction := NewMovementAction(-1, 0)
		return &movementAction
	case ebiten.KeyArrowRight, ebiten.KeyD:
		movementAction := NewMovementAction(1, 0)
		return &movementAction
	case ebiten.KeyEscape:
		return EscapeAction{}
	}
	return nil
}
