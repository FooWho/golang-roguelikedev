package main

import (
	_ "image/png"
	"log"

	"github.com/FooWho/golang-roguelikedev/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	engine := &engine.Engine{}
	engine.Initialize(80, 50, 800, 500, 10)

	ebiten.SetWindowSize(engine.GetSize())
	ebiten.SetWindowTitle("Golang RogueLikeDev Tutorial")

	if err := ebiten.RunGame(engine); err != nil {
		log.Fatal(err)
	}

}
