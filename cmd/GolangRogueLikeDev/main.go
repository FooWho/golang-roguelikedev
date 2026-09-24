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
	tiles   []*ebiten.Image
}

func (g *Game) Update() error {
	action := g.EventHandler()
	if action == nil {
		return nil
	}
	if action != nil {
		switch v := action.(type) {
		case *MovementAction:
			g.playerX += v.dx
			g.playerY += v.dy
		case *EscapeAction:
			return ebiten.Termination
		default:
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	char := '@'
	tileSprite := g.tiles[char]

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
		tiles:   loadTileset(),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Golang RogueLikeDev Tutorial")

	if err := ebiten.RunGame(game); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}

}

func loadTileset() []*ebiten.Image {
	file, err := os.Open("/home/jelison/Workspace/GolangRogueLikeDev/assets/dejavu10x10_gs_tc.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
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
