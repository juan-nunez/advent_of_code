package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    f, _ := os.Open("input.txt")
    defer f.Close()

    scanner := bufio.NewScanner(f)

    freqChanges := make([]int, 0)
    freqMap := make(map[int]int)

    for scanner.Scan() {
        change, _ := strconv.Atoi(scanner.Text())
        freqChanges = append(freqChanges, change)
    }

    var frequency int = 0

    for {
        for i := 0; i < len(freqChanges); i++ {
            frequency += freqChanges[i]
            freqMap[frequency]++
            if freqMap[frequency] == 2 {
                fmt.Println(frequency)
                return
            }
        }
    }
}

