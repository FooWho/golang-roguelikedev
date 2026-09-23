package engine

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func (e *Engine) Draw(screen *ebiten.Image) {

	tilesPerRow := 32
	char := '@'
	charIndex := int(char) - 32

	spriteX := (charIndex % tilesPerRow) * e.tileSize
	spriteY := (charIndex / tilesPerRow) * e.tileSize

	rect := image.Rect(spriteX, spriteY, spriteX+e.tileSize, spriteY+e.tileSize)

	tileSprite := e.tileSet.SubImage(rect).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.playerX*e.tileSize), float64(e.playerY*e.tileSize))

	screen.DrawImage(tileSprite, op)
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

func (e *Engine) MoveEntity(dx int, dy int) {
	newX := e.playerX + dx
	newY := e.playerY + dy
	if e.IsOnScreen(newX, newY) {
		e.playerX += dx
		e.playerY += dy
	}
}

func loadTileset() *ebiten.Image {
	file, err := os.Open("/home/jelison/Workspace/GolangRogueLikeDev/assets/dejavu10x10_gs_tc.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func (e *Engine) GetSize() (int, int) {
	return e.screenWidth, e.screenHeight
}
