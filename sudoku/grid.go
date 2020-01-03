package sudoku

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Grid struct {
	// Cells are indexed by x and y. x increases from left to right. y
	// increases from top to bottom. (0, 0) is the top left, (8, 8) is the
	// bottom right. To get cell (x, y), use Grid.Cells[y][x]
	cells [][]*Cell
}

func NewGrid() *Grid {
	grid := &Grid{}
	grid.cells = make([][]*Cell, 9)
	return grid
}

func (g *Grid) String() string {
	var b bytes.Buffer
	for y, row := range g.cells {
		for x, cell := range row {
			switch x {
			case 3, 6:
				b.WriteString("|")
			}
			if cell.Value == 0 {
				b.WriteString(".")
			} else {
				b.WriteString(fmt.Sprint(cell.Value))
			}
		}
		switch y {
		case 2, 5:
			b.WriteString("\n-----------")
		}
		b.WriteString("\n")
	}
	s := b.String()
	s = strings.TrimSpace(s)
	return s
}

func (g *Grid) Copy() *Grid {
	s := g.String()
	s = strings.ReplaceAll(s, ".", "0")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "|", "")
	s = strings.ReplaceAll(s, "\n\n", "\n")

	new := NewGrid()
	if err := new.Parse(s); err != nil {
		panic(err)
	}
	return new
}

// TODO:
// - validation
// - ignore grid lines in input
func (g *Grid) Parse(s string) error {
	s = strings.TrimSpace(s)
	for y, values := range strings.Split(s, "\n") {
		row := make([]*Cell, 9)
		for x, val := range values {
			intVal, err := strconv.Atoi(string(val))
			if err != nil {
				return err
			}
			row[x] = NewCell(intVal, x, y)
		}
		g.cells[y] = row
	}
	return nil
}

func (g *Grid) Cell(x, y int) *Cell {
	if x < 0 || x >= 9 || y < 0 || y >= 9 {
		log.Fatalf("cannot get cell (%d, %d)", x, y)
	}
	return g.cells[y][x]
}

func (g *Grid) SetCell(x, y, value int) {
	logger.Debug("setting (%d, %d) to %d", x, y, value)
	cell := g.Cell(x, y)
	cell.Value = value
}

func (g *Grid) Cols() [][]*Cell {
	cols := make([][]*Cell, 9)
	for i := 0; i < 9; i++ {
		cols[i] = make([]*Cell, 9)
	}
	for y, row := range g.cells {
		for x, cell := range row {
			cols[x][y] = cell
		}
	}
	return cols
}

func (g *Grid) Col(x int) []*Cell {
	return g.Cols()[x]
}

func (g *Grid) Row(y int) []*Cell {
	return g.cells[y]
}

// Returns the cells in the square that contains (x, y)
func (g *Grid) Square(x, y int) []*Cell {
	minMax := func(i int) (min int, max int) {
		if i < 0 {
			log.Fatalf("i less than 0: %d", i)
		}
		if i < 3 {
			return 0, 3
		}
		if i < 6 {
			return 3, 6
		}
		if i < 9 {
			return 6, 9
		}
		log.Fatalf("i over 9: %d", i)
		return 0, 0
	}

	xmin, xmax := minMax(x)
	ymin, ymax := minMax(y)

	var cells []*Cell
	for y1 := ymin; y1 < ymax; y1++ {
		for x1 := xmin; x1 < xmax; x1++ {
			cells = append(cells, g.Cell(x1, y1))
		}
	}
	return cells
}

func (g *Grid) Squares() [][]*Cell {
	squares := make([][]*Cell, 0, 9)
	for yoffset := 0; yoffset <= 6; yoffset += 3 {
		for xoffset := 0; xoffset <= 6; xoffset += 3 {
			var square []*Cell
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					x1 := x + xoffset
					y1 := y + yoffset
					cell := g.Cell(x1, y1)
					square = append(square, cell)
				}
			}
			squares = append(squares, square)
		}
	}
	return squares
}

func (g *Grid) Solved() bool {
	// TODO: validate grid is correctly solved
	for _, row := range g.cells {
		for _, cell := range row {
			if cell.Value == 0 {
				return false
			}
		}
	}
	return true
}
