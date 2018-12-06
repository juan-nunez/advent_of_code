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

    var frequency int = 0
    for scanner.Scan() {
        change, _ := strconv.Atoi(scanner.Text())
        frequency += change
    }

    fmt.Println(frequency)
}
