package engine

type Tile struct {
	visual      Visual
	walkable    bool
	transparent bool
}

func NewTile(visual Visual, walkable bool, transparent bool) *Tile {
	return &Tile{
		visual:      visual,
		walkable:    walkable,
		transparent: transparent,
	}
}

type GameMap struct {
	width  int
	height int
	tiles  []Tile
}

func NewGameMap(width int, height int) *GameMap {
	return &GameMap{width: width, height: height, tiles: make([]Tile, width*height, width*height)}
}

func (gm *GameMap) Fill() {
	tile := NewTile(NewVisual(' ', Color{red: 0, green: 0, blue: 0}, Color{255, 255, 255}), true, true)

	for y := 0; y < gm.height; y++ {
		for x := 0; x < gm.width; x++ {
			gm.tiles[y*gm.width+x] = *tile
		}
	}
}

func (gm *GameMap) InBounds(x int, y int) bool {
	return x >= 0 && x < gm.width && y >= 0 && y < gm.height
}
