package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		os.Exit(1)
	}

	width, height := screen.Size()
	game := NewGame(width, height)

	dirChan := make(chan Point)
	quitChan := make(chan bool)

	go func() {
		for {
			select {
			case <-quitChan:
				return
			default:
				ev := screen.PollEvent()
				switch ev := ev.(type) {
				case *tcell.EventKey:
					switch ev.Key() {
					case tcell.KeyUp:
						dirChan <- Point{X: 0, Y: -1}
					case tcell.KeyDown:
						dirChan <- Point{X: 0, Y: 1}
					case tcell.KeyLeft:
						dirChan <- Point{X: -1, Y: 0}
					case tcell.KeyRight:
						dirChan <- Point{X: 1, Y: 0}
					case tcell.KeyEscape:
						os.Exit(0)
					}
				}
			}
		}
	}()

	for {
		select {
		case dir := <-dirChan:
			game.Snake.Direction = dir
		default:
		}

		if !game.MoveSnake() {
			break
		}

		DrawGame(screen, &game)
		time.Sleep(150 * time.Millisecond)
	}

	quitChan <- true

	currHigh := LoadHighScore()
	if game.Score > currHigh {
		SaveHighScore(game.Score)
	}
	game.Highscore = LoadHighScore()

	DrawGameOver(screen, &game)

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				screen.Fini()
				os.Exit(0)
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'q', 'Q':
					screen.Fini()
					os.Exit(0)
				case 'r', 'R':
					screen.Fini()
					main()
				}
			}
		}
	}
}
