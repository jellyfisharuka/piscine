package main

import (
	"os"

	"github.com/01-edu/z01"
)

type Table [9][9]int

var solutionCount int

func main() {
	ok := checkInput(os.Args[1:])
	if !ok {
		println("Error")
		return
	}
	table, ok := loadTable()
	if !ok {
		println("Error")
		return
	}

	numSolutions := countSolutions(table)
	if numSolutions != 1 {
		println("Error")
		return
	}

	ok = solve(table)
	if !ok {
		println("Error")
		return
	}
	printTable(table)
}

func countSolutions(t *Table) int {
	solutionCount = 0
	solveWithCount(t)
	return solutionCount
}

func solveWithCount(t *Table) {
	for i := 0; i < 9; i++ {
		row := t[i]
		for j := 0; j < 9; j++ {
			colVal := row[j]
			if colVal == 0 {
				for n := 1; n <= 9; n++ {
					if valid(t, i, j, n) {
						t[i][j] = n
						solveWithCount(t)
						t[i][j] = 0

						// Check if two solutions are found and return early
						if solutionCount >= 2 {
							return
						}
					}
				}
				return
			}
		}
	}
	solutionCount++
}

func checkInput(args []string) bool {
	if len(args) != 9 {
		return false
	}

	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			return false
		}
	}

	for i := 0; i < len(args); i++ {
		for _, value := range args[i] {
			if value == '/' || value == '0' {
				return false
			} else if value < '.' || value > '9' {
				return false
			}
		}
	}

	return true
}

func println(s string) {
	print(s)
	z01.PrintRune('\n')
}

func print(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func loadTable() (*Table, bool) {
	t := &Table{}

	rows := os.Args[1:]

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n, ok := atoi(rows[i][j])
			if !ok {
				return t, false
			}

			if n != 0 {
				if !valid(t, i, j, n) {
					return t, false
				}
			}

			t[i][j] = n
		}
	}

	return t, true
}

func atoi(r byte) (int, bool) {
	c := rune(r)

	if c == '.' {
		return 0, true
	}

	if c < '0' || c > '9' {
		return -1, false
	}

	return int(c - '0'), true
}

func valid(t *Table, row int, col int, n int) bool {
	for i := 0; i < 9; i++ {
		if i != col && t[row][i] == n {
			return false
		}

		if i != row && t[i][col] == n {
			return false
		}
	}

	rowStart, colStart := row/3*3, col/3*3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if i != row && j != col && t[i][j] == n {
				return false
			}
		}
	}

	return true
}

func printInt(i int) {
	r := '0' + i

	z01.PrintRune(rune(r))
}

func printTable(t *Table) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 8; j++ {
			printInt(t[i][j])
			print(" ")
		}
		printInt(t[i][8])
		println("")
	}
}

func solve(t *Table) bool {
	for i := 0; i < 9; i++ {
		row := t[i]
		for j := 0; j < 9; j++ {
			colVal := row[j]
			if colVal == 0 {
				for n := 1; n <= 9; n++ {
					if valid(t, i, j, n) {
						t[i][j] = n
						if solve(t) {
							return true
						}
						t[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}
