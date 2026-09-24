package engine

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/FooWho/golang-roguelikedev/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

func (e *Engine) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	for _, renderable := range e.renderables {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(renderable.GetX()*e.tileSize), float64(renderable.GetY()*e.tileSize))
		tileSprite := e.tiles[renderable.GetVisual().char]
		screen.DrawImage(tileSprite, op)
	}
}

func (e *Engine) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return e.screenWidth, e.screenHeight
}

func (e *Engine) IsOnScreen(x int, y int) bool {
	if x >= 0 && x < e.gridWidth && y >= 0 && y < e.gridHeight {
		return true
	}
	return false
}

func loadTileset(tileSize int) []*ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(assets.TilesetData))
	if err != nil {
		log.Fatal(err)
	}

	spriteSheet := ebiten.NewImageFromImage(img)
	tiles := make([]*ebiten.Image, 256)
	tilesPerRow := 32

	for ascii := 32; ascii < 256; ascii++ {
		sheetIndex := ascii - 32
		spriteX := (sheetIndex % tilesPerRow) * tileSize
		spriteY := (sheetIndex / tilesPerRow) * tileSize
		rect := image.Rect(spriteX, spriteY, spriteX+tileSize, spriteY+tileSize)

		tiles[ascii] = spriteSheet.SubImage(rect).(*ebiten.Image)
	}

	return tiles
}

func (e *Engine) GetSize() (int, int) {
	return e.screenWidth, e.screenHeight
}
