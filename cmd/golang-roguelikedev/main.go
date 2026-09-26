package main

import (
	_ "embed"
	"encoding/json"
	_ "image/png"
	"log"

	"github.com/FooWho/golang-roguelikedev/assets"
	"github.com/FooWho/golang-roguelikedev/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed properties.json
var propsFile []byte

type Properties struct {
	GridWidth  map[string]int `json:"gridWidth"`
	GridHeight map[string]int `json:"gridHeight"`
	TileSize   map[string]int `json:"tileSize`
}

func main() {
	var properties Properties
	err := json.Unmarshal(propsFile, &properties)
	if err != nil {
		log.Fatal("Could not load/unmarshall properties.json!")
	}
	var artExists string
	var artType int
	if assets.FileExists(assets.AssetsFS, "Full.png") {
		log.Print("Loaded Art")
		artExists = "demonic_dungeon"
		artType = 0
	} else if assets.FileExists(assets.AssetsFS, "dejavu10x10_gs+tc.png") {
		log.Print("Loaded ASCII")
		artExists = "ascii"
		artType = 1
	} else {
		log.Fatal("Neither ASCIItileset nor Art tilesets exist!")
	}
	engine := engine.NewEngine(properties.GridWidth[artExists],
		properties.GridHeight[artExists],
		properties.TileSize[artExists],
		artType,
	)

	ebiten.SetWindowSize(engine.GetSize())
	ebiten.SetWindowTitle("Golang RogueLikeDev Tutorial")

	if err := ebiten.RunGame(engine); err != nil {
		if err != ebiten.Termination {
			log.Fatal(err)
		}

	}

}
