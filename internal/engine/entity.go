package engine

type Entity struct {
	x      int
	y      int
	visual Visual
}

func NewEntity(x int, y int, visual Visual) Entity {
	return Entity{x: x, y: y, visual: visual}
}

func (e *Entity) GetX() int {
	return e.x
}

func (e *Entity) GetY() int {
	return e.y
}

func (e *Entity) GetVisual() Visual {
	return e.visual
}

// Interface Guard
var _ Renderable = (*Entity)(nil)

type Actor interface {
	GetAction(engine *Engine) Action
	//GetEntity() *Entity
}

type Monster struct {
	*Entity
}

func (m *Monster) GetAction(engine *Engine) Action {
	return (Action)(nil)
}

/*
func (m *Monster) GetEntity() *Entity {
	return m.Entity
}*/

type Player struct {
	*Entity
}

func NewPlayer(e *Entity) *Player {
	return &Player{Entity: e}
}

func (p *Player) GetAction(engine *Engine) Action {
	return engine.GetPlayerAction()
}

/*
func (p *Player) GetEntity() *Entity {
	return p.Entity
}*/

type Renderable interface {
	GetX() int
	GetY() int
	GetVisual() Visual
}

type Visual struct {
	char  rune
	color Color
}

type Color struct {
	red   int
	blue  int
	green int
}

func NewColor(r int, g int, b int) Color {
	return Color{red: r, green: g, blue: b}
}

func NewVisual(char rune, color Color) Visual {
	return Visual{char: char, color: color}
}
