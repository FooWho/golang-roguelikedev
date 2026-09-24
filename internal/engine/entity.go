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

func (e *Entity) SetPosition(x int, y int) {
	e.x = x
	e.y = y
}

func (e *Entity) GetVisual() Visual {
	return e.visual
}

// Interface Guard
var _ Renderable = (*Entity)(nil)

type Actor interface {
	GetAction(engine *Engine) Action
	GetX() int
	GetY() int
	SetPosition(x int, y int)
}

type Monster struct {
	*Entity
}

func NewMonster(e *Entity) *Monster {
	return &Monster{Entity: e}
}

func (m *Monster) GetAction(engine *Engine) Action {
	return NewWaitAction()
}

// Interface Guard
var _ Actor = (*Monster)(nil)
var _ Renderable = (*Monster)(nil)

type Player struct {
	*Entity
}

func NewPlayer(e *Entity) *Player {
	return &Player{Entity: e}
}

func (p *Player) GetAction(engine *Engine) Action {
	return engine.GetPlayerAction()
}

// Interface Guard
var _ Actor = (*Player)(nil)
var _ Renderable = (*Player)(nil)

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
