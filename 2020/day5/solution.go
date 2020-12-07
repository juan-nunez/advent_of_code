package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func readInput() []string {
	f, _ := os.Open("input.txt")
	defer f.Close()

	boardingPasses := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		boardingPasses = append(boardingPasses, scanner.Text())
	}

	return boardingPasses
}

func binarySpacePartition(s string, length, idx, low, high int) int {
	mid := low + (high-low)/2
	if idx == length-1 {
		if s[idx] == 'F' || s[idx] == 'L' {
			return low
		} else {
			return high
		}
	} else if s[idx] == 'F' || s[idx] == 'L' {
		return binarySpacePartition(s, length, idx+1, low, mid)
	} else if s[idx] == 'B' || s[idx] == 'R' {
		return binarySpacePartition(s, length, idx+1, mid+1, high)
	}

	return -1
}

func getSeat(boardingPass string) (int, int) {
	row := binarySpacePartition(boardingPass[:7], 7, 0, 0, 127)
	col := binarySpacePartition(boardingPass[7:], 3, 0, 0, 7)
	return row, col
}

func getSeatId(boardingPass string) int {
	col, row := getSeat(boardingPass)
	return col*8 + row
}

func getHighestSeatId(boardingPasses []string) int {
	max := 0
	for _, boardingPass := range boardingPasses {
		max = int(math.Max(float64(getSeatId(boardingPass)), float64(max)))
	}
	return max
}

func doPartOne() int {
	boardingPasses := readInput()
	highestSeatId := getHighestSeatId(boardingPasses)
	return highestSeatId
}

func doPartTwo() int {
	boardingPasses := readInput()
	boardingPassIds := make([]int, 0)
	for _, boardingPass := range boardingPasses {
		boardingPassIds = append(boardingPassIds, getSeatId(boardingPass))
	}

	sort.Ints(boardingPassIds)
	mySeatId := -1
	for i := 1; i < len(boardingPassIds)-1; i++ {
		if boardingPassIds[i-1] != boardingPassIds[i]-1 {
			mySeatId = boardingPassIds[i] - 1
			break
		}
	}

	return mySeatId
}

func main() {
	fmt.Println(doPartOne())
	fmt.Println(doPartTwo())
}
