package main

import (
	"math/rand"
	"time"
)

func NewGame(width, height int) Game {
	startPos := Point{X: width / 2, Y: height / 2}

	snake := Snake{
		Body: []Point{
			startPos,
			{X: startPos.X - 1, Y: startPos.Y},
			{X: startPos.X - 2, Y: startPos.Y},
		},
		Direction: Point{X: 1, Y: 0},
	}

	game := Game{
		Snake:     snake,
		Width:     width,
		Height:    height,
		Score:     0,
		Highscore: LoadHighScore(),
	}

	rand.Seed(time.Now().UnixNano())
	game.SpawnFood()
	return game
}

func (g *Game) SpawnFood() {
	g.Food = Food{
		Position: Point{
			X: rand.Intn(g.Width),
			Y: rand.Intn(g.Height),
		},
	}
}

func (g *Game) MoveSnake() bool {
	head := g.Snake.Body[0]

	newHead := Point{
		X: head.X + g.Snake.Direction.X,
		Y: head.Y + g.Snake.Direction.Y,
	}

	if newHead.X < 0 || newHead.X >= g.Width ||
		newHead.Y < 0 || newHead.Y >= g.Height {
		return false
	}

	for _, segment := range g.Snake.Body {
		if segment == newHead {
			return false
		}
	}

	if newHead == g.Food.Position {
		g.Score++
		g.Snake.Body = append([]Point{newHead}, g.Snake.Body...)
		g.SpawnFood()
	} else {
		g.Snake.Body = append([]Point{newHead}, g.Snake.Body[:len(g.Snake.Body)-1]...)
	}

	return true
}
