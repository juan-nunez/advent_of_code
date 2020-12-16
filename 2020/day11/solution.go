package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func readLayout() []string {
	f, _ := os.Open("input.txt")
	defer f.Close()

	layout := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		layout = append(layout, scanner.Text())
	}

	return layout
}

func isAdjacentOccupied(layout []string, i, j int) bool {
	maxi := len(layout)
	maxj := len(layout[0])

	isOccupiedAdjacent := false
	for _, dir := range dirs {
		nexti := i + dir[0]
		nextj := j + dir[1]
		if nexti < 0 || nexti >= maxi || nextj < 0 || nextj >= maxj {
			continue
		}

		if layout[nexti][nextj] == '#' {
			isOccupiedAdjacent = true
			break
		}
	}

	return isOccupiedAdjacent
}

func shouldVacate(layout []string, i, j int) bool {
	maxi := len(layout)
	maxj := len(layout[0])

	numOccupied := 0
	for _, dir := range dirs {
		nexti := i + dir[0]
		nextj := j + dir[1]
		if nexti < 0 || nexti >= maxi || nextj < 0 || nextj >= maxj {
			continue
		}

		if layout[nexti][nextj] == '#' {
			numOccupied++
		}
	}

	return numOccupied >= 4
}

func isAdjacentOccupiedPartTwo(layout []string, i, j int) bool {
	maxi := len(layout)
	maxj := len(layout[0])

	for _, dir := range dirs {
		nexti := i
		nextj := j
		for {
			nexti += dir[0]
			nextj += dir[1]
			if nexti < 0 || nexti >= maxi || nextj < 0 || nextj >= maxj {
				break
			}

			if layout[nexti][nextj] == 'L' {
				break
			}

			if layout[nexti][nextj] == '#' {
				return true
			}
		}
	}

	return false
}

func shouldVacatePartTwo(layout []string, i, j int) bool {
	maxi := len(layout)
	maxj := len(layout[0])
	numOccupied := 0
	for _, dir := range dirs {
		nexti := i
		nextj := j
		for {
			nexti += dir[0]
			nextj += dir[1]
			if nexti < 0 || nexti >= maxi || nextj < 0 || nextj >= maxj {
				break
			}

			if layout[nexti][nextj] == 'L' {
				break
			}

			if layout[nexti][nextj] == '#' {
				numOccupied++
				break
			}
		}
	}

	return numOccupied >= 5
}

func doPartOne() {
	layout := readLayout()
	for {
		changed := false
		occupiedSeats := 0
		layoutCopy := make([]string, len(layout))
		copy(layoutCopy, layout)
		for i := 0; i < len(layout); i++ {
			for j := 0; j < len(layout[0]); j++ {
				if layout[i][j] == '#' {
					occupiedSeats++
				}
				if layout[i][j] == '#' && shouldVacate(layout, i, j) {
					layoutCopy[i] = layoutCopy[i][0:j] + "L" + layoutCopy[i][j+1:]
					changed = true
				} else if layout[i][j] == 'L' && !isAdjacentOccupied(layout, i, j) {
					layoutCopy[i] = layoutCopy[i][0:j] + "#" + layoutCopy[i][j+1:]
					changed = true
				}
			}
		}
		layout = layoutCopy
		if !changed {
			fmt.Println(occupiedSeats)
			break
		}
	}
}

func doPartTwo() {
	layout := readLayout()
	for {
		changed := false
		occupiedSeats := 0
		layoutCopy := make([]string, len(layout))
		copy(layoutCopy, layout)
		for i := 0; i < len(layout); i++ {
			for j := 0; j < len(layout[0]); j++ {
				if layout[i][j] == '#' {
					occupiedSeats++
				}
				if layout[i][j] == '#' && shouldVacatePartTwo(layout, i, j) {
					layoutCopy[i] = layoutCopy[i][0:j] + "L" + layoutCopy[i][j+1:]
					changed = true
				} else if layout[i][j] == 'L' && !isAdjacentOccupiedPartTwo(layout, i, j) {
					layoutCopy[i] = layoutCopy[i][0:j] + "#" + layoutCopy[i][j+1:]
					changed = true
				}
			}
		}
		layout = layoutCopy
		if !changed {
			fmt.Println(occupiedSeats)
			break
		}
	}

}

func main() {
	doPartOne()
	doPartTwo()
}
