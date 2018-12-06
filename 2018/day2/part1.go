package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    f, _ := os.Open("input.txt")
    defer f.Close()

    scanner := bufio.NewScanner(f)

    numTwos := 0
    numThrees := 0

    for scanner.Scan() {
        word := scanner.Text()
        twosFound := false
        threesFound := false
        letterMap := make(map[rune]int)
        for _, letter := range word {
            letterMap[letter]++
        }

        for _, v := range letterMap {
            if !twosFound && v == 2 {
                numTwos++
                twosFound = true
            }

            if !threesFound && v == 3 {
                numThrees++
                threesFound = true
            }
        }
    }

    fmt.Println(numTwos * numThrees)

}
