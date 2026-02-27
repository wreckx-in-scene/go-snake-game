package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func DrawGame(screen tcell.Screen, game *Game) {
	screen.Clear()

	//styles
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorWheat)
	foodStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)
	headStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	bodyStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow)
	scoreStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)

	//drawing border
	for x := 0; x < game.Width; x++ {
		screen.SetContent(x, 0, '#', nil, borderStyle)
		screen.SetContent(x, game.Height-1, '#', nil, borderStyle)
	}

	for y := 0; y < game.Height; y++ {
		screen.SetContent(0, y, '#', nil, borderStyle)
		screen.SetContent(game.Width-1, y, '#', nil, borderStyle)
	}

	//drawing food
	screen.SetContent(
		game.Food.Position.X,
		game.Food.Position.Y,
		'*', nil, foodStyle)

	//drawing snake
	for i, segment := range game.Snake.Body {
		if i == 0 {
			screen.SetContent(segment.X, segment.Y, '@', nil, headStyle)
		} else {
			screen.SetContent(segment.X, segment.Y, 'o', nil, bodyStyle)
		}
	}

	//draw score
	scoreText := fmt.Sprintf("Score: %d  High: %d", game.Score, game.Highscore)
	for i, ch := range scoreText {
		screen.SetContent(i, game.Height-1, ch, nil, scoreStyle)
	}

	screen.Show()

}

func DrawText(screen tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, ch := range text {
		screen.SetContent(x+i, y, ch, nil, style)
	}
}

func DrawGameOver(screen tcell.Screen, game *Game) {
	screen.Clear()

	style := tcell.StyleDefault.Foreground(tcell.ColorRed)
	scoreStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	highStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow)

	centerX := game.Width / 2
	centerY := game.Height / 2

	//drawing game over screen
	DrawText(screen, centerX-5, centerY-3, "Game Over!", style)
	DrawText(screen, centerX-5, centerY-1, fmt.Sprintf("Score : %d", game.Score), scoreStyle)
	DrawText(screen, centerX-5, centerY, fmt.Sprintf("High Score : %d", game.Highscore), highStyle)
	DrawText(screen, centerX-5, centerY+2, "R - Restart", scoreStyle)
	DrawText(screen, centerX-5, centerY+3, "Q - Quit", scoreStyle)
}
