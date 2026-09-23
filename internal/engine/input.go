package engine

import (
	"github.com/FooWho/golang-roguelikedev/internal/actions"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (e *Engine) EventHandler() (actions.Action, error) {
	keys := make([]ebiten.Key, 0, 5)
	keys = inpututil.AppendJustPressedKeys(keys)

	if len(keys) == 0 {
		return nil, nil
	}

	switch keys[0] {
	case ebiten.KeyArrowUp, ebiten.KeyW:
		return actions.NewMovementAction(0, -1), nil
	case ebiten.KeyArrowDown, ebiten.KeyS:
		return actions.NewMovementAction(0, 1), nil
	case ebiten.KeyArrowLeft, ebiten.KeyA:
		return actions.NewMovementAction(-1, 0), nil
	case ebiten.KeyArrowRight, ebiten.KeyD:
		return actions.NewMovementAction(1, 0), nil
	}
	return nil, nil
}
