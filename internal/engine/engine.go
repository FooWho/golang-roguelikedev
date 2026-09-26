package engine

type Engine struct {
	tiles
	gridWidth    int
	gridHeight   int
	screenWidth  int
	screenHeight int
	tileSize     int
	player       *Player
	actors       []Actor
	renderables  []Renderable
	gameMap      *GameMap
	artType      int
}

func NewEngine(gridWidth int, gridHeight int, tileSize int, artType int) *Engine {
	tiles := loadTileset(tileSize)

	pe := NewEntity(gridWidth/2, gridHeight/2, NewVisual("hero"))
	player := NewPlayer(&pe)

	actors := make([]Actor, 0, 100)
	actors = append(actors, player)

	renderables := make([]Renderable, 0, 100)
	renderables = append(renderables, player)

	gm := NewGameMap(gridWidth, gridHeight)
	gm.Fill()

	return &Engine{
		gridWidth:    gridWidth,
		gridHeight:   gridHeight,
		screenWidth:  gridWidth * tileSize,
		screenHeight: gridHeight * tileSize,
		tileSize:     tileSize,
		tiles:        *S,
		player:       player,
		actors:       actors,
		renderables:  renderables,
		gameMap:      gm,
	}
}

func (e *Engine) Update() error {
	playerAction := e.player.GetAction(e)

	if playerAction == nil {
		return nil
	}
	err := playerAction.Perform(e.player, e)
	if err != nil {
		return err
	}

	for i := 1; i < len(e.actors); i++ {
		monster := e.actors[i]
		monsterAction := monster.GetAction(e)
		err = monsterAction.Perform(monster, e)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) GetPlayer() *Player {
	return e.player
}
