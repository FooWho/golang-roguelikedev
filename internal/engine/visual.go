package engine

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
