package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func readGrid() []string {
	f, _ := os.Open("input.txt")
	defer f.Close()

	grid := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	return grid
}

func extend(grid []string, yslope int) []string {
	length := len(grid)
	width := len(grid[0])

	requiredWidth := length * yslope
	// if width is already biggger than the required width, then return the existing grid.
	if width >= requiredWidth {
		return grid
	}
	newGrid := make([]string, length)
	copy(newGrid, grid)

	numCopy := int(math.Ceil(float64(requiredWidth) / float64(width)))

	for i := 0; i < numCopy; i++ {
		for j := 0; j < length; j++ {
			newGrid[j] += grid[j]
		}
	}

	return newGrid
}

func treesFound(xslope, yslope int) int {
	grid := readGrid()
	grid = extend(grid, yslope)

	numTrees := 0

	length := len(grid)
	width := len(grid[0])

	i, j := 0, 0
	for i < length && j < width {
		val := grid[i][j]
		if string(val) == "#" {
			numTrees++
		}

		i += xslope
		j += yslope
	}

	return numTrees
}

func main() {
	// Part 1 uses a slope of 1,3
	fmt.Println(treesFound(1, 3))

	// Part 2 uses various slopes
	slopes := [][2]int{
		{1, 3},
		{1, 1},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	ans := 1
	for _, slope := range slopes {
		found := treesFound(slope[0], slope[1])
		ans *= found
	}
	fmt.Println(ans)
}
