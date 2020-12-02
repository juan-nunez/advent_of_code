package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type PasswordPayload struct {
	low      int
	high     int
	letter   rune
	password string
}

func readInput() []PasswordPayload {
	f, _ := os.Open("input.txt")
	defer f.Close()

	input := make([]PasswordPayload, 0)
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		parts := re.FindStringSubmatch(text)
		low, _ := strconv.Atoi(parts[1])
		high, _ := strconv.Atoi(parts[2])
		letter := rune(parts[3][0])
		password := parts[4]
		passwordPayload := PasswordPayload{low: low, high: high, letter: letter, password: password}
		input = append(input, passwordPayload)
	}

	return input
}

func part1IsValid(low, high int, letter rune, password string) bool {
	letterMap := make(map[rune]int, 0)
	for _, letter := range password {
		letterMap[letter]++
	}

	if letterMap[letter] < low || letterMap[letter] > high {
		return false
	}

	return true
}

func doPart1() int {
	passwordPayloads := readInput()
	validPasswordCnt := 0
	for _, p := range passwordPayloads {
		if part1IsValid(p.low, p.high, p.letter, p.password) {
			validPasswordCnt++
		}
	}
	return validPasswordCnt
}

func part2IsValid(low, high int, letter rune, password string) bool {
	// If letter is in one index AND letter is not in both indices
	if (rune(password[low-1]) == letter || rune(password[high-1]) == letter) &&
		!(rune(password[low-1]) == letter && rune(password[high-1]) == letter) {
		return true
	}

	return false
}

func doPart2() int {
	passwordPayloads := readInput()
	validPasswordCnt := 0
	for _, p := range passwordPayloads {
		if part2IsValid(p.low, p.high, p.letter, p.password) {
			validPasswordCnt++
		}
	}
	return validPasswordCnt
}

func main() {
	fmt.Println(doPart1())
	fmt.Println(doPart2())
}
