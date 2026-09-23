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

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.playerY > 0 {
			g.playerY--
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.playerY < gridHeight-1 {
			g.playerY++
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.playerX > 0 {
			g.playerX--
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.playerX < gridWidth-1 {
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

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var tileset *ebiten.Image

func main() {

	game := &Game{
		playerX: int(gridWidth / 2),
		playerY: int(gridHeight / 2),
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
