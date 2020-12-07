package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() [][]string {
	f, _ := os.Open("input.txt")
	defer f.Close()

	groups := make([][]string, 0)
	scanner := bufio.NewScanner(f)
	groupAnswers := make([]string, 0)
	for scanner.Scan() {
		if scanner.Text() != "" {
			groupAnswers = append(groupAnswers, scanner.Text())
			continue
		}

		groups = append(groups, groupAnswers)
		groupAnswers = make([]string, 0)
	}

	// Handle last group
	groups = append(groups, groupAnswers)

	return groups
}

func countUnique(s string) int {
	m := make(map[rune]bool, 0)
	cnt := 0

	for _, letter := range s {
		if !m[letter] {
			m[letter] = true
			cnt++
		}
	}

	return cnt
}

func countCommon(group []string) int {
	cnt := 0
	hashmap := make(map[rune]int, 0)
	set := make(map[rune]bool, 0)
	for _, answers := range group {
		for _, letter := range answers {
			if !set[letter] {
				set[letter] = true
				hashmap[letter]++
			}
		}
		set = make(map[rune]bool, 0)
	}

	for _, v := range hashmap {
		if v == len(group) {
			cnt++
		}
	}

	return cnt
}

func doPartOne() int {
	cnt := 0
	groups := readInput()
	for _, group := range groups {
		combinedGroupAnswer := ""
		for _, groupAnswer := range group {
			combinedGroupAnswer += groupAnswer
		}
		cnt += countUnique(combinedGroupAnswer)
	}

	return cnt
}

func doPartTwo() int {
	cnt := 0
	groups := readInput()
	for _, group := range groups {
		cnt += countCommon(group)
	}

	return cnt
}

func main() {
	fmt.Println(doPartOne())
	fmt.Println(doPartTwo())
}
