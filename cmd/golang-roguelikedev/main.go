package main

import (
	_ "image/png"
	"log"

	"github.com/FooWho/golang-roguelikedev/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gridWidth    = 80
	gridHeight   = 50
	tilesize     = 16
	screenWidth  = gridWidth * tilesize
	screenHeight = gridHeight * tilesize
)

func main() {

	engine := engine.NewEngine(gridWidth, gridHeight, screenWidth, screenHeight, tilesize)

	ebiten.SetWindowSize(engine.GetSize())
	ebiten.SetWindowTitle("Golang RogueLikeDev Tutorial")

	if err := ebiten.RunGame(engine); err != nil {
		if err != ebiten.Termination {
			log.Fatal(err)
		}

	}

}
