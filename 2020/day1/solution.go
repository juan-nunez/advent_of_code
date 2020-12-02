package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() []int {
	f, _ := os.Open("input.txt")
	defer f.Close()

	nums := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	return nums
}

func doPart1() int {
	nums := readInput()
	target := 2020
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return nums[i] * nums[j]
			}
		}
	}

	return -1
}

func doPart2() int {
	nums := readInput()
	target := 2020
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == target {
					return nums[i] * nums[j] * nums[k]
				}
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(doPart1())
	fmt.Println(doPart2())
}
