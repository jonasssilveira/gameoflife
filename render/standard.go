package render

import (
	"conway-game-of-life/types"
	"fmt"
)

func RenderBoardStandard(board types.Board) {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for p := range board {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			if board[types.Point{X: x, Y: y}] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
