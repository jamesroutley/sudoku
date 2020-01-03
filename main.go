package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jamesroutley/sudoku/sudoku"
)

func main() {
	puzzleBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	grid := sudoku.NewGrid()
	if err := grid.Parse(string(puzzleBytes)); err != nil {
		log.Fatal(err)
	}
	if err := sudoku.Solve(grid); err != nil {
		fmt.Println(grid)
		log.Fatal(err)
	}

	fmt.Println("Solved!")
	fmt.Printf("%s\n", grid)
}
