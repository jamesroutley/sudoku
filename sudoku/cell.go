package sudoku

import (
	"fmt"
)

type Cell struct {
	Value int
	x     int
	y     int
}

func NewCell(val, x, y int) *Cell {
	return &Cell{
		Value: val,
		x:     x,
		y:     y,
	}
}

func (c *Cell) String() string {
	return fmt.Sprint(c.Value)
}

func (c *Cell) getPossibleValues(g *Grid) ([]int, error) {
	if c.Value != 0 {
		return nil, nil
	}
	possibilities := map[int]bool{}
	for i := 1; i <= 9; i++ {
		possibilities[i] = true
	}

	impossibilities := g.Row(c.y)
	impossibilities = append(impossibilities, g.Col(c.x)...)
	impossibilities = append(impossibilities, g.Square(c.x, c.y)...)

	for _, c := range impossibilities {
		if c.Value == 0 {
			continue
		}
		delete(possibilities, c.Value)
	}

	if len(possibilities) == 0 {
		return nil, fmt.Errorf("No possibilities for (%d, %d)", c.x, c.y)
	}
	return intSetToSlice(possibilities), nil
}
