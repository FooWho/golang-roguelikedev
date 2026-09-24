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
	entities     []Entity
}

func NewEngine(gridWidth int, gridHeight int, screenWidth int, screenHeight int, tileSize int) Engine {
	tiles := loadTileset(tileSize)

	player := NewEntity(gridWidth/2, gridHeight/2, NewVisual('@', NewColor(255, 255, 255)))
	entities := make([]Entity, 0, 10)
	entities = append(entities, player)

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
		entities:     entities,
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
