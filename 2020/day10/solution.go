package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readJoltages() []int {
	f, _ := os.Open("input.txt")
	defer f.Close()

	joltages := make([]int, 0)
	joltages = append(joltages, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		joltage, _ := strconv.Atoi(scanner.Text())
		joltages = append(joltages, joltage)
	}

	sort.Ints(joltages)
	max := joltages[len(joltages)-1] + 3
	joltages = append(joltages, max)
	return joltages
}

func doPartOne() {
	joltages := readJoltages()

	numOnes := 0
	numThrees := 0

	for i := 0; i < len(joltages)-1; i++ {
		if joltages[i+1]-joltages[i] == 1 {
			numOnes++
		} else if joltages[i+1]-joltages[i] == 3 {
			numThrees++
		}
	}

	fmt.Println(numOnes * numThrees)
}

func countPaths(m map[int]bool, memo map[int]int64, node int, target int) int64 {
	if node == target {
		return 1
	}

	if memo[node] != -1 {
		return memo[node]
	}

	var paths int64 = 0
	for i := 1; i <= 3; i++ {
		// If not in list
		if !m[node+i] {
			continue
		}

		paths += countPaths(m, memo, node+i, target)
	}

	memo[node] = paths
	return paths
}

func doPartTwo() {
	joltages := readJoltages()

	m := make(map[int]bool, 0)
	memo := make(map[int]int64, 0)

	for _, joltage := range joltages {
		m[joltage] = true
		memo[joltage] = -1
	}

	target := joltages[len(joltages)-1]
	numPaths := countPaths(m, memo, 0, target)
	fmt.Println(numPaths)
}

func main() {
	doPartOne()
	doPartTwo()
}
