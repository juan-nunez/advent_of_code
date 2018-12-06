package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    f, _ := os.Open("input.txt")
    defer f.Close()

    scanner := bufio.NewScanner(f)
    words := make([]string, 0)
    for scanner.Scan() {
        word := scanner.Text()
        words = append(words, word)
    }

    for i := 0; i < len(words) - 1; i++ {
        cWord := words[i]
        for j := i+1; j < len(words); j++ {
            comparedWord := words[j]
            differAmt := 0
            var commonLetters strings.Builder
            for k := 0; k < len(cWord) && differAmt < 2; k++ {
                if cWord[k] == comparedWord[k] {
                    commonLetters.WriteString(string(cWord[k]))
                } else {
                    differAmt++
                }
            }

            if differAmt < 2 {
                fmt.Println(commonLetters.String())
                return
            }
        }
    }
}
