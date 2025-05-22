package render

import (
	"conway-game-of-life/types"
	"os"

	"golang.org/x/term"
)

type Renderer func(types.Board)

func GetRenderer() Renderer {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		return RenderBoardTerminal
	}
	return RenderBoardStandard
}
