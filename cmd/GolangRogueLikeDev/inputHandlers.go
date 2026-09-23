package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) EventHandler() (Action, error) {
	keys := make([]ebiten.Key, 0, 5)
	keys = inpututil.AppendJustPressedKeys(keys)

	if len(keys) == 0 {
		return nil, nil
	}

	switch keys[0] {
	case ebiten.KeyArrowUp, ebiten.KeyW:
		return &MovementAction{dx: 0, dy: -1}, nil
	case ebiten.KeyArrowDown, ebiten.KeyS:
		return &MovementAction{dx: 0, dy: 1}, nil
	case ebiten.KeyArrowLeft, ebiten.KeyA:
		return &MovementAction{dx: -1, dy: 0}, nil
	case ebiten.KeyArrowRight, ebiten.KeyD:
		return &MovementAction{dx: 1, dy: 0}, nil
	}
	return nil, nil
}
