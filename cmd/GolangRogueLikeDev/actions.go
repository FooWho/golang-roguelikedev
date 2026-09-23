package main

import "fmt"

type Action interface {
	Perform(g *Game)
}

type MovementAction struct {
	dx int
	dy int
}

func (m *MovementAction) Perform(g *Game) {
	if g.playerX+m.dx >= 0 && g.playerX+m.dx < gridWidth {
		g.playerX += m.dx
	}
	if g.playerY+m.dy >= 0 && g.playerY+m.dy < gridHeight {
		g.playerY += m.dy
	}
	fmt.Printf("{g.playerX: %d, g.playerY: %d}\n", g.playerX, g.playerY)
}
