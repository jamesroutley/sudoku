# Sudoku solver

Reads a sudoku puzzle in from stdin, and prints out the solution

```sh
sudoku % go run main.go << EOL
008001000
020806050
000007004
972000010
000070000
060000837
500300000
080102060
000700500
EOL
Solved!
648|531|972
729|846|351
153|927|684
-----------
972|683|415
831|475|296
465|219|837
-----------
597|364|128
384|152|769
216|798|543
```

The code is a _real mess_
