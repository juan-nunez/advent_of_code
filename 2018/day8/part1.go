package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func walkTreeRecursive(input []int, idx int) (int, int) {
    numChildren := input[idx]
    numMetadata := input[idx+1]
    idx += 2
    metadataValue := 0
    var value int
    for i := 0; i < numChildren; i++ {
        idx, value = buildTreeRecursive(input, idx)
        metadataValue += value
    }

    for i := 0; i < numMetadata; i++ {
        metadataValue += input[idx]
        idx++
    }
    return idx, metadataValue
}

func main() {
    content, _ := ioutil.ReadFile("input.txt")

    input := make([]int, 0)
    for _, in := range strings.Split(strings.TrimSpace(string(content[:len(content)])), " ") {
        num, _ := strconv.Atoi(in)
        input = append(input, num)
    }

    _, metadataValue := walkTreeRecursive(input, 0)
    fmt.Println(metadataValue)
}
