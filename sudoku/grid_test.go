package sudoku

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSquare(t *testing.T) {
	grid := NewGrid()

	err := grid.Parse(`
003020600
900305001
001806400
008102900
700000008
006708200
002609500
800203009
005010300`)
	require.NoError(t, err)
	square := grid.Square(1, 1)
	for _, cell := range square {
		// fmt.Println(cell.Value)
		_ = cell
	}
	// assert.True(t, false)
}
