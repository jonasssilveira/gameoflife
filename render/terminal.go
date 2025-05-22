package render

import (
	"conway-game-of-life/types"
	"fmt"
	"golang.org/x/term"
	"os"
)

func RenderBoardTerminal(board types.Board) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Could not get terminal size:", err)
		return
	}

	centerX := width / 2
	centerY := height / 2

	screen := make([][]rune, height)
	for i := range screen {
		screen[i] = make([]rune, width)
		for j := range screen[i] {
			screen[i][j] = ' '
		}
	}

	for point := range board {
		sx := centerX + int(point.X)
		sy := centerY - int(point.Y)

		if sx >= 0 && sx < width && sy >= 0 && sy < height {
			screen[sy][sx] = '█'
		}
	}
	for _, row := range screen {
		fmt.Println(string(row))
	}
}
