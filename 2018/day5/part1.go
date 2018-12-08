package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func doesReact(a, b string) bool {
    capital := strings.ToUpper(a)
    lower := strings.ToLower(a)
    if (a == capital && b == lower) || (a == lower && b == capital) {
        return true
    }
    return false
}

func main() {
    f, _ := os.Open("input.txt")
    defer f.Close()
    scanner := bufio.NewScanner(f)
    var polymer string;

    for scanner.Scan() {
        polymer += scanner.Text()
    }

    polymer += "."
    for {
        var i int = 0
        length := len(polymer)
        var deleted bool = false;
        var polymerReacted string

        for i < length-1 {
            if doesReact(string(polymer[i]), string(polymer[i+1])) {
                deleted = true
                i+=2 
            } else {
                polymerReacted += string(polymer[i])
                i++
            }
        }

        polymer = polymerReacted
        polymer += "."

        if !deleted {
            fmt.Println(len(polymer)-1)
            return
        }
    }
}
