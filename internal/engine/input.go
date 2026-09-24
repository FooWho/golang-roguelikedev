package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (e *Engine) GetPlayerAction() Action {
	e.keys = inpututil.AppendJustPressedKeys(e.keys[:0])

	if len(e.keys) == 0 {
		return nil
	}

	switch e.keys[0] {
	case ebiten.KeyArrowUp, ebiten.KeyW:
		return NewMovementAction(0, -1)
	case ebiten.KeyArrowDown, ebiten.KeyS:
		return NewMovementAction(0, 1)
	case ebiten.KeyArrowLeft, ebiten.KeyA:
		return NewMovementAction(-1, 0)
	case ebiten.KeyArrowRight, ebiten.KeyD:
		return NewMovementAction(1, 0)
	case ebiten.KeyEscape:
		return NewEscapeAction()
	}
	return nil
}
