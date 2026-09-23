package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
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
	tileSet *ebiten.Image
}

func (g *Game) Update() error {
	a, err := g.EventHandler()
	if err != nil {
		log.Fatal(err)
	}
	if a != nil {
		a.Perform(g)
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

	tileSprite := g.tileSet.SubImage(rect).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.playerX*tileSize), float64(g.playerY*tileSize))

	screen.DrawImage(tileSprite, op)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	game := &Game{
		playerX: int(gridWidth / 2),
		playerY: int(gridHeight / 2),
		tileSet: loadTileset(),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Golang RogueLikeDev Tutorial")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
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
