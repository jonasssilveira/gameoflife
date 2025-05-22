package life

import (
	"conway-game-of-life/types"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestNextGen_Blinker_Success(t *testing.T) {
	initial := types.Board{
		{1, 0}: true,
		{1, 1}: true,
		{1, 2}: true,
	}
	expected := types.Board{
		{0, 1}: true,
		{1, 1}: true,
		{2, 1}: true,
	}

	result := NextGen(initial)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNextGen_Blinker_Fail(t *testing.T) {
	initial := types.Board{
		{1, 0}: true,
		{1, 1}: true,
		{1, 2}: true,
	}
	expected := types.Board{
		{0, 1}: true,
		{1, 1}: true,
		{2, 1}: true,
	}

	result := NextGen(initial)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestLoadLife106(t *testing.T) {
	content := `
#Life 1.06
0 0
1 0
1 1
	`
	expected := types.Board{
		{0, 0}: true,
		{1, 0}: true,
		{1, 1}: true,
	}

	board := parseLife106String(content)

	if !reflect.DeepEqual(board, expected) {
		t.Errorf("Expected %v, got %v", expected, board)
	}
}

func parseLife106String(input string) types.Board {
	board := make(types.Board)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0 {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		board[types.Point{X: x, Y: y}] = true
	}
	return board
}
