package life

import (
	"bufio"
	"conway-game-of-life/types"
	"os"
	"strconv"
	"strings"
)

func LoadLife106(path string) (types.Board, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	board := make(types.Board)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		x, err1 := strconv.Atoi(parts[0])
		y, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		board[types.Point{X: x, Y: y}] = true
	}

	return board, nil
}
