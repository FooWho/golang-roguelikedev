package engine

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"log"

	"github.com/FooWho/golang-roguelikedev/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

func (e *Engine) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// 1. Draw Map Tiles
	for y := 0; y < e.gameMap.height; y++ {
		for x := 0; x < e.gameMap.width; x++ {
			tile := e.gameMap.tiles[e.gameMap.GetIndex(x, y)]

			op.GeoM.Reset()
			op.GeoM.Translate(float64(x*e.tileSize), float64(y*e.tileSize))

			// Fetch the exact image using the string key!
			spriteImg := e.spriteSheet.Sprites[tile.visual.spriteName]
			screen.DrawImage(spriteImg, op)
		}
	}

	// 2. Draw Renderables
	for _, renderable := range e.renderables {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(renderable.GetX()*e.tileSize), float64(renderable.GetY()*e.tileSize))

		spriteImg := e.spriteSheet.Sprites[renderable.GetVisual().spriteName]
		screen.DrawImage(spriteImg, op)
	}
}
func (e *Engine) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return e.screenWidth, e.screenHeight
}

func (e *Engine) IsOnScreen(x int, y int) bool {
	return x >= 0 && x < e.gridWidth && y >= 0 && y < e.gridHeight
}

func NewVisual(name string) Visual {
	return Visual{spriteName: name}
}

type Visual struct {
	spriteName string
}

func NewSpriteSheet(image *ebiten.Image, sprites map[string]*ebiten.Image) *SpriteSheet {
	return &SpriteSheet{Image: image, Sprites: sprites}
}

type SpriteSheet struct {
	Image   *ebiten.Image
	Sprites map[string]*ebiten.Image
}

func loadTileset(tileSize int) *SpriteSheet {
	imgData := assets.GetSpriteSheet()
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Fatal(err)
	}

	sheet := ebiten.NewImageFromImage(img)
	sprites := make(map[string]*ebiten.Image)

	extractSprite := func(name string, gridX, gridY int) {
		pixelX := gridX * tileSize
		pixelY := gridY * tileSize
		fmt.Printf("Loading %s at (%d, %d)\n", name, pixelX, pixelY)
		rect := image.Rect(pixelX, pixelY, pixelX+tileSize, pixelY+tileSize)
		sprites[name] = sheet.SubImage(rect).(*ebiten.Image)
	}

	// Map Tiles
	extractSprite("wall", 0, 0)  // Example: Top left corner wall block
	extractSprite("floor", 8, 0) // Example: The blank dark floor tile

	// Entities
	extractSprite("player", 26, 0)   // Example: The first human frame
	extractSprite("skeleton", 16, 0) // Example: The first skeleton frame
	extractSprite("slime", 18, 4)    // Example: The green slime frame

	return &SpriteSheet{
		Image:   sheet,
		Sprites: sprites,
	}
}

func (e *Engine) GetSize() (int, int) {
	return e.screenWidth, e.screenHeight
}
