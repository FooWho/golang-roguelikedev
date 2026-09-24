package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Engine struct {
	playerX      int
	playerY      int
	tiles        []*ebiten.Image
	gridWidth    int
	gridHeight   int
	screenWidth  int
	screenHeight int
	tileSize     int
	keys         []ebiten.Key
}

func NewEngine(gridWidth int, gridHeight int, screenWidth int, screenHeight int, tileSize int) Engine {
	tiles := loadTileset(tileSize)

	return Engine{
		gridWidth:    gridWidth,
		gridHeight:   gridHeight,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		tileSize:     tileSize,
		tiles:        tiles,
		playerX:      gridWidth / 2,
		playerY:      gridHeight / 2,
		keys:         make([]ebiten.Key, 0, 5),
	}
}

func (e *Engine) Update() error {
	action := e.GetAction()

	if action == nil {
		return nil
	} else {
		switch v := action.(type) {
		case *MovementAction:
			v.Perform(e)
		case EscapeAction:
			return ebiten.Termination
		}
	}

	return nil
}
