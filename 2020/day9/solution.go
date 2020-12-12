package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func readInput() []int64 {
	f, _ := os.Open("input.txt")
	defer f.Close()
	input := make([]int64, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		input = append(input, num)
	}

	return input
}

func doPartOne() int64 {
	nums := readInput()
	mp := make(map[int64]int, 0)

	// Add first 25 numbers to map
	for i := 0; i < 25; i++ {
		mp[nums[i]]++
	}

	// Iterate through entire input
	for i := 25; i < len(nums); i++ {
		curr := nums[i]
		found := false
		// Iterate through last 25 numbers
		for j := i - 25; j < i; j++ {
			// If curr - a number in the previous 25 is in the map
			if mp[curr-nums[j]] > 0 {
				// Result of curr-nums[j] == nums[j] is only valid IF there are multiple of that number
				if curr-nums[j] == nums[j] && mp[nums[j]] < 2 {
					continue
				}
				found = true
				break
			}
		}
		if !found {
			return curr
		}

		// Slide the window one index over. Remove i-25th index and add curr
		mp[nums[i-25]]--
		mp[nums[i]]++
	}
	return -1
}

func doPartTwo(invalidNumber int64) int64 {
	nums := readInput()
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i]
		mn := int64(math.Min(float64(nums[i]), 9223372036854775807))
		mx := int64(math.Max(float64(nums[i]), 0))
		for j := i + 1; j < len(nums); j++ {
			mn = int64(math.Min(float64(nums[j]), float64(mn)))
			mx = int64(math.Max(float64(nums[j]), float64(mx)))
			sum += nums[j]
			if sum == invalidNumber {
				return mn + mx
			}
		}
	}
	return -1
}

func main() {
	invalidNumber := doPartOne()
	fmt.Println(invalidNumber)
	fmt.Println(doPartTwo(invalidNumber))
}
