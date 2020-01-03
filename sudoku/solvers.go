package sudoku

import (
	"fmt"
	"sort"
)

type Solver func(*Grid) (solved bool, err error)

type CellSolver func(g *Grid, x, y int) (solved bool, err error)

var solvers = [...]Solver{
	constructSolver(UniquePossibilityCellSolver),
	constructSolver(OnlyFitsHereCellSolver),
}

func Solve(g *Grid) error {
	for !g.Solved() {
		progressMade, err := runSolvers(g)
		if err != nil {
			return err
		}
		if !progressMade {
			return fmt.Errorf("Can't solve any more")
		}
	}
	return nil
}

func runSolvers(g *Grid) (progressMade bool, err error) {
	// Run non-guessing solvers first
	for _, solver := range solvers {
		progressMade, err := solver(g)
		if err != nil {
			return false, err
		}
		if progressMade {
			return true, nil
		}
	}
	// Haven't made any progess without guessing - let's try that
	progressMade, err = guessSolver(g)
	if err != nil {
		return false, err
	}
	if progressMade {
		return true, nil
	}
	return false, nil
}

func constructSolver(cs CellSolver) Solver {
	return func(g *Grid) (solved bool, err error) {
		for y := 0; y < 9; y++ {
			for x := 0; x < 9; x++ {
				solved, err := cs(g, x, y)
				if err != nil {
					return false, err
				}
				if solved {
					return true, nil
				}
			}
		}
		return false, nil
	}
}

func intSetToSlice(m map[int]bool) []int {
	var result []int
	for i := range m {
		result = append(result, i)
	}
	sort.Ints(result)
	return result
}

func intSliceToSet(s []int) map[int]bool {
	result := map[int]bool{}
	for _, i := range s {
		result[i] = true
	}
	return result
}
