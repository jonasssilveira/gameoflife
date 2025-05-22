package life

import (
	"conway-game-of-life/types"
)

var directions = []types.Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func NextGen(current types.Board) types.Board {
	neighborCounts := make(map[types.Point]int)
	for cell := range current {
		for _, d := range directions {
			neighbor := types.Point{X: cell.X + d.X, Y: cell.Y + d.Y}
			neighborCounts[neighbor]++
		}
	}

	next := make(types.Board)

	for cell, count := range neighborCounts {
		if count == 3 || (count == 2 && current[cell]) {
			next[cell] = true
		}
	}

	return next
}
