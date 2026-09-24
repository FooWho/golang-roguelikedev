package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func GetAction() (Action, error) {
	keys := make([]ebiten.Key, 0, 5)
	keys = inpututil.AppendJustPressedKeys(keys)

	if len(keys) == 0 {
		return nil, nil
	}

	switch keys[0] {
	case ebiten.KeyArrowUp, ebiten.KeyW:
		return NewMovementAction(0, -1), nil
	case ebiten.KeyArrowDown, ebiten.KeyS:
		return NewMovementAction(0, 1), nil
	case ebiten.KeyArrowLeft, ebiten.KeyA:
		return NewMovementAction(-1, 0), nil
	case ebiten.KeyArrowRight, ebiten.KeyD:
		return NewMovementAction(1, 0), nil
	}
	return nil, nil
}
