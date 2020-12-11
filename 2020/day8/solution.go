package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	name string
	arg  int
}

func readBootcode() []instruction {
	f, _ := os.Open("input.txt")
	defer f.Close()

	re := regexp.MustCompile(`([a-z]+) (\+|-)([0-9]+)`)
	instructions := make([]instruction, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		instruction := instruction{}
		parts := re.FindStringSubmatch(scanner.Text())
		instruction.name = parts[1]
		sign := parts[2]
		magnitude, _ := strconv.Atoi(parts[3])
		if sign == "-" {
			magnitude = -magnitude
		}
		instruction.arg = magnitude

		instructions = append(instructions, instruction)
	}

	return instructions
}

func doPartOne() int {
	instructions := readBootcode()
	visited := make([]bool, len(instructions))
	acc := 0
	pc := 0

	for {
		if visited[pc] {
			break
		}

		visited[pc] = true
		instruction := instructions[pc]
		switch instruction.name {
		case "acc":
			acc += instruction.arg
			pc++
		case "jmp":
			pc += instruction.arg
		case "nop":
			pc++
		default:
			fmt.Println("Instruction not found")
		}
	}

	return acc
}

func doPartTwo() int {
	instructions := readBootcode()
	acc := 0
	found := false
	for i := 0; i < len(instructions) && !found; i++ {
		oldInstr := instructions[i]
		newInstr := instruction{name: "nop"}
		instructions[i] = newInstr
		visited := make([]bool, len(instructions))
		pc := 0
		acc = 0
		for {
			if pc >= len(instructions) {
				found = true
				break
			}

			if visited[pc] {
				break
			}

			visited[pc] = true
			instruction := instructions[pc]
			switch instruction.name {
			case "acc":
				acc += instruction.arg
				pc++
			case "jmp":
				pc += instruction.arg
			case "nop":
				pc++
			default:
				fmt.Println("Instruction not found")
			}
		}

		instructions[i] = oldInstr
	}

	return acc
}

func main() {
	//fmt.Println(doPartOne())
	fmt.Println(doPartTwo())
}
