package life

import (
	"conway-game-of-life/types"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadLife106Table(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		fileContent string
		want        types.Board
		shouldErr   bool
	}{
		{
			name: "valid file",
			fileContent: `
						#Life 1.06
						0 0
						1 0
						1 1
						`,
			want: types.Board{
				{0, 0}: true,
				{1, 0}: true,
				{1, 1}: true,
			},
			shouldErr: false,
		},
		{
			name:        "file not found",
			fileContent: "",
			want:        nil,
			shouldErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var path string
			if tt.name != "file not found" {
				path = filepath.Join(tmpDir, "life.txt")
				err := os.WriteFile(path, []byte(tt.fileContent), 0644)
				if err != nil {
					t.Fatalf("failed to write temp file: %v", err)
				}
			} else {
				path = filepath.Join(tmpDir, "nonexistent.txt")
			}

			board, err := LoadLife106(path)

			if tt.shouldErr && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(board, tt.want) {
				t.Errorf("expected board %v, got %v", tt.want, board)
			}
		})
	}
}
