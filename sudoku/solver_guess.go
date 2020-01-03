package sudoku

import "log"

var guessesDeep = 0

// Find a square with two possibilities and guess at one. If the guess is a
// failure, try the other option
func guessSolver(g *Grid) (solved bool, err error) {
	guessesDeep += 1
	defer func() {
		logger.Debug("Leaving %d guess solver", guessesDeep)
		guessesDeep -= 1
	}()
	logger.Debug("cannot solve deterministically - guessing. %d guesses deep", guessesDeep)

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			cell := g.Cell(x, y)
			possibilities, err := cell.getPossibleValues(g)
			if err != nil {
				logger.Debug("Error finding possibilities for cell (%d, %d) in the guessSolver", x, y)
				return false, err
			}
			if len(possibilities) != 2 {
				continue
			}

			for i, possibility := range possibilities {
				g1 := g.Copy()
				logger.Debug("solving (%d, %d) attempt %d by guessing possibility to be: %d", x, y, i+1, possibility)
				g1.SetCell(x, y, possibility)
				err := Solve(g1)
				if err != nil {
					logger.Debug("attempt %d failed", i+1)
					// Only return the error if we're on the last iteration
					if i == (len(possibilities) - 1) {
						return false, err
					}
					// Otherwise, continue searching
					continue
				} else {
					// We've solved it. Fill in the original g and return
					// TODO: do better
					for y := 0; y < 9; y++ {
						for x := 0; x < 9; x++ {
							g.SetCell(x, y, g1.Cell(x, y).Value)
						}
					}
					return true, nil
				}
			}
		}
	}
	log.Fatalf("could not find cell with two possibilities")
	return false, nil
}
