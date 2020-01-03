package sudoku

// UniquePossibilityCellSolver solves squares which only have a single possible
// value
func UniquePossibilityCellSolver(g *Grid, x, y int) (solved bool, err error) {
	cell := g.Cell(x, y)
	if cell.Value != 0 {
		return false, nil
	}
	possibilities, err := cell.getPossibleValues(g)
	if err != nil {
		return false, err
	}

	if len(possibilities) == 1 {
		g.SetCell(cell.x, cell.y, possibilities[0])
		return true, nil
	}
	return false, nil
}
