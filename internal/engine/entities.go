package engine

type Entity struct {
	x      int
	y      int
	visual Visual
}

func NewEntity(x int, y int, visual Visual) Entity {
	return Entity{x: x, y: y, visual: visual}
}
