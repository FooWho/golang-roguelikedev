package engine

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"log"

	"github.com/FooWho/golang-roguelikedev/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

var whitePixel *ebiten.Image

func init() {
	whitePixel = ebiten.NewImage(1, 1)
	whitePixel.Fill(color.White)
}

func (e *Engine) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	for y := 0; y < e.gridHeight; y++ {
		for x := 0; x < e.gridWidth; x++ {
			tile := e.gameMap.tiles[y*e.gridWidth+x]
			vis := tile.visual

			op.GeoM.Reset()
			op.GeoM.Scale(float64(e.tileSize), float64(e.tileSize))
			op.GeoM.Translate(float64(x*e.tileSize), float64(y*e.tileSize))

			op.ColorScale.Reset()
			//op.ColorScale.Scale(float32(vis.bg.red)/255.0, float32(vis.bg.green)/255.0, float32(vis.bg.blue)/255.0, 1)

			screen.DrawImage(whitePixel, op)

			op.GeoM.Reset()
			op.GeoM.Translate(float64(x*e.tileSize), float64(y*e.tileSize))

			op.ColorScale.Reset()
			//op.ColorScale.Scale(float32(vis.fg.red)/255.0, float32(vis.fg.green)/255.0, float32(vis.fg.blue)/255.0, 1)

			screen.DrawImage(e.tiles[vis.spriteName], op)
		}
	}

	for _, renderable := range e.renderables {
		vis := renderable.GetVisual()

		op.GeoM.Reset()
		op.GeoM.Scale(float64(e.tileSize), float64(e.tileSize))
		op.GeoM.Translate(float64(renderable.GetX()*e.tileSize), float64(renderable.GetY()*e.tileSize))

		op.ColorScale.Reset()
		op.ColorScale.Scale(
			float32(vis.bg.red)/255.0,
			float32(vis.bg.green)/255.0,
			float32(vis.bg.blue)/255.0,
			1,
		)
		screen.DrawImage(whitePixel, op)

		op.GeoM.Reset()
		op.GeoM.Translate(float64(renderable.GetX()*e.tileSize), float64(renderable.GetY()*e.tileSize))

		op.ColorScale.Reset()
		op.ColorScale.Scale(
			float32(vis.fg.red)/255.0,
			float32(vis.fg.green)/255.0,
			float32(vis.fg.blue)/255.0,
			1,
		)
		charIndex := vis.char
		if charIndex < 0 || int(charIndex) >= len(e.tiles) {
			charIndex = '?'
		}
		tileSprite := e.tiles[charIndex]
		screen.DrawImage(tileSprite, op)
	}
}

func (e *Engine) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return e.screenWidth, e.screenHeight
}

func (e *Engine) IsOnScreen(x int, y int) bool {
	return x >= 0 && x < e.gridWidth && y >= 0 && y < e.gridHeight
}

type Visual struct {
	spriteName string
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
		rect := image.Rect(pixelX, pixelY, pixelX+tileSize, pixelY+tileSize)
		sprites[name] = sheet.SubImage(rect).(*ebiten.Image)
	}

	// Map Tiles
	extractSprite("wall", 0, 0)  // Example: Top left corner wall block
	extractSprite("floor", 0, 8) // Example: The blank dark floor tile

	// Entities
	extractSprite("player", 0, 24)   // Example: The first human frame
	extractSprite("skeleton", 0, 16) // Example: The first skeleton frame
	extractSprite("slime", 4, 18)    // Example: The green slime frame

	return &SpriteSheet{
		Image:   sheet,
		Sprites: sprites,
	}
}

func (e *Engine) GetSize() (int, int) {
	return e.screenWidth, e.screenHeight
}
