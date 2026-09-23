package engine

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Engine struct {
	playerX      int
	playerY      int
	tileSet      *ebiten.Image
	gridWidth    int
	gridHeight   int
	screenWidth  int
	screenHeight int
	tileSize     int
}

func NewEngine(gridWidth int, gridHeight int, screenWidth int, screenHeight int, tileSize int) Engine {
	return Engine{
		gridWidth:    gridWidth,
		gridHeight:   gridHeight,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		tileSize:     tileSize,
		tileSet:      loadTileset(),
		playerX:      gridWidth / 2,
		playerY:      gridHeight / 2,
	}
}

func (e *Engine) Update() error {
	a, err := e.EventHandler()
	if err != nil {
		log.Fatal(err)
	}
	if a != nil {
		a.Perform(e)
	}

	return nil
}
