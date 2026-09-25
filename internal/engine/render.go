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
			op.ColorScale.Scale(float32(vis.bg.red)/255.0, float32(vis.bg.green)/255.0, float32(vis.bg.blue)/255.0, 1)

			screen.DrawImage(whitePixel, op)

			op.GeoM.Reset()
			op.GeoM.Translate(float64(x*e.tileSize), float64(y*e.tileSize))

			op.ColorScale.Reset()
			op.ColorScale.Scale(float32(vis.fg.red)/255.0, float32(vis.fg.green)/255.0, float32(vis.fg.blue)/255.0, 1)

			screen.DrawImage(e.tiles[vis.char], op)
		}
	}

	for _, renderable := range e.renderables {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(renderable.GetX()*e.tileSize), float64(renderable.GetY()*e.tileSize))

		vis := renderable.GetVisual()

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

func loadTileset(tileSize int) []*ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(assets.TilesetData))
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	transparentImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			r, g, b, _ := pixelColor.RGBA()

			// If the pixel is pure black (0, 0, 0)
			if r == 0 && g == 0 && b == 0 {
				// Make it fully transparent
				transparentImg.Set(x, y, color.Transparent)
			} else {
				// Keep the original white character pixel
				transparentImg.Set(x, y, pixelColor)
			}
		}
	}

	// Use our new transparent image instead of the original
	spriteSheet := ebiten.NewImageFromImage(transparentImg)
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
