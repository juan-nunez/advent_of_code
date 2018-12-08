package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

var (
    alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func fullyReact(polymer string) int {
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
            return len(polymer)-1
        }
    }
}

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

    minLength := 99999999
    for i := 0; i < 1 /*len(alphabet)*/; i++ {
        newPolymer := ""
        for j := 0; j < len(polymer); j++ {
            if string(polymer[j]) != string(alphabet[i]) && string(polymer[j]) != string(strings.ToUpper(string(alphabet[i]))) {
                newPolymer += string(polymer[j])
            }
        }
        newPolymer += "."
        length := fullyReact(newPolymer)
        if length < minLength {
            minLength = length
        }
    }

    fmt.Println(minLength)

}
