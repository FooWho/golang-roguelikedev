package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	gridWidth    = 80
	gridHeight   = 50
	screenWidth  = gridWidth * tileSize
	screenHeight = gridHeight * tileSize
	tileSize     = 10
)

type Game struct {
	playerX int
	playerY int
}

func (g *Game) Update() error {
	// Move Up
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.playerY > 0 { // Prevent moving off the top edge
			g.playerY--
		}
	}
	// Move Down
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.playerY < gridHeight-1 { // Prevent moving off the bottom edge
			g.playerY++
		}
	}
	// Move Left
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.playerX > 0 { // Prevent moving off the left edge
			g.playerX--
		}
	}
	// Move Right
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.playerX < gridWidth-1 { // Prevent moving off the right edge
			g.playerX++
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	tilesPerRow := 32
	char := '@'
	charIndex := int(char) - 32

	spriteX := (charIndex % tilesPerRow) * tileSize
	spriteY := (charIndex / tilesPerRow) * tileSize

	rect := image.Rect(spriteX, spriteY, spriteX+tileSize, spriteY+tileSize)

	tileSprite := tileset.SubImage(rect).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.playerX*tileSize), float64(g.playerY*tileSize))

	screen.DrawImage(tileSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var tileset *ebiten.Image

func main() {

	game := &Game{
		playerX: 40,
		playerY: 25,
	}

	loadTileset()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Golang RogueLikeDev Tutorial")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}

func loadTileset() {
	file, err := os.Open("assets/dejavu10x10_gs_tc.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	tileset = ebiten.NewImageFromImage(img)
}
