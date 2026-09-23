package engine

import (
	"image"
	"log"
	"os"

	"github.com/FooWho/golang-roguelikedev/internal/actions"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Engine struct {
	playerX      int
	playerY      int
	tileSet      *ebiten.Image
	gridWidth    int
	gridHeight   int
	screenWidth  int
	screenHeight int
	tileSize     int
}

func (e *Engine) Initialize(
	gridWidth int,
	gridHeight int,
	screenWidth int,
	screenHeight int,
	tileSize int,
) {
	e.gridWidth = gridWidth
	e.gridHeight = gridHeight
	e.screenWidth = screenWidth
	e.screenHeight = screenHeight
	e.tileSize = tileSize
	e.tileSet = loadTileset()
	e.playerX = (gridWidth / 2)
	e.playerY = (gridHeight / 2)
}

func (e *Engine) GetSize() (int, int) {
	return e.screenWidth, e.screenHeight
}

func (e *Engine) Update() error {
	a, err := e.EventHandler()
	if err != nil {
		log.Fatal(err)
	}
	if a != nil {
		a.Perform(e)
	}

	return nil
}

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

func (e *Engine) EventHandler() (actions.Action, error) {
	keys := make([]ebiten.Key, 0, 5)
	keys = inpututil.AppendJustPressedKeys(keys)

	if len(keys) == 0 {
		return nil, nil
	}

	switch keys[0] {
	case ebiten.KeyArrowUp, ebiten.KeyW:
		return actions.NewMovementAction(0, -1), nil
	case ebiten.KeyArrowDown, ebiten.KeyS:
		return actions.NewMovementAction(0, 1), nil
	case ebiten.KeyArrowLeft, ebiten.KeyA:
		return actions.NewMovementAction(-1, 0), nil
	case ebiten.KeyArrowRight, ebiten.KeyD:
		return actions.NewMovementAction(1, 0), nil
	}
	return nil, nil
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
