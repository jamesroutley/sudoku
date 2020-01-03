package sudoku

import "fmt"

func OnlyFitsHereCellSolver(g *Grid, x, y int) (solved bool, err error) {
	cell := g.Cell(x, y)
	if cell.Value != 0 {
		return false, nil
	}

	// TODO: rename this
	solveGroup := func(cells []*Cell) (solved bool, err error) {
		onlyFitsHereSlice, err := cell.getPossibleValues(g)
		if err != nil {
			return false, err
		}
		onlyFitsHere := intSliceToSet(onlyFitsHereSlice)

		for _, c := range cells {
			// Don't look at the cell we're currently considering
			if c.x == x && c.y == y {
				continue
			}
			possibilities, err := c.getPossibleValues(g)
			if err != nil {
				return false, err
			}
			for _, possibility := range possibilities {
				delete(onlyFitsHere, possibility)
			}
		}

		if len(onlyFitsHere) > 1 {
			return false, fmt.Errorf("more than one number fits at (%d, %d): %+v", x, y, onlyFitsHere)
		}

		if len(onlyFitsHere) == 1 {
			for possibility := range onlyFitsHere {
				g.SetCell(cell.x, cell.y, possibility)
				return true, nil
			}
		}
		return false, nil
	}

	solved, err = solveGroup(g.Row(cell.y))
	if err != nil {
		return false, err
	}
	if solved {
		return true, nil
	}

	solved, err = solveGroup(g.Col(cell.x))
	if err != nil {
		return false, err
	}
	if solved {
		return true, nil
	}
	solved, err = solveGroup(g.Square(cell.x, cell.y))

	if err != nil {
		return false, err
	}
	if solved {
		return true, nil
	}

	return false, nil
}
