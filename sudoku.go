package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const SIZE = 9

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		line := make([]int, SIZE)
		for idx, r := range lineStr {
			if r != '.' {
				n, er := strconv.Atoi(string(r))
				if er != nil {
					panic(fmt.Sprintf("Not a number %c", n))
				}
				line[idx] = n
			}
		}
		grid = append(grid, line)
	}
	fmt.Println(grid)
	start := time.Now()
	solve(grid)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time =", elapsed)
}

func solve(grid [][]int) {
	unsolved := true
	if grid[0][0] > 0 {
		panic("The first cell must be blank")
	}
	for n := 1; n < SIZE+1; n++ {
		ok, g := addNumber(n, 0, 0, grid)
		if ok {
			fmt.Println(g)
			unsolved = false
			break
		}
	}
	if unsolved {
		fmt.Println("No solution found!")
	}
}

func addNumber(n, row, col int, grid [][]int) (bool, [][]int) {
	// If number fits, add another number
	// Return true if finished
	if rowOK(n, row, grid) && colOK(n, col, grid) && boxOk(n, col, row, grid) {
		gridCopy := make([][]int, SIZE)
		for idx, r := range grid {
			rc := make([]int, SIZE)
			copy(rc, r)
			gridCopy[idx] = rc
		}
		gridCopy[row][col] = n
		for {
			col++
			if col == SIZE {
				col = 0
				row++
				if row == SIZE {
					return true, gridCopy // Solved
				}
			}
			if gridCopy[row][col] == 0 {
				for i := 1; i < SIZE+1; i++ {
					ok, g := addNumber(i, row, col, gridCopy)
					if ok {
						return true, g
					}
				}
				break // Reached a dead end
			}
		}
	}
	return false, grid
}

func rowOK(n, row int, grid [][]int) bool {
	for _, m := range grid[row] {
		if n == m {
			return false
		}
	}
	return true
}

func colOK(n, col int, grid [][]int) bool {
	for row := 0; row < len(grid); row++ {
		if n == grid[row][col] {
			return false
		}
	}
	return true
}

func boxOk(n, col, row int, grid [][]int) bool {
	x0 := col / 3 * 3
	y0 := row / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if n == grid[y0+i][x0+j] {
				return false
			}
		}
	}
	return true
}
