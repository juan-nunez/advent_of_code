package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)


func buildTreeRecursive(input []int, idx int) (int, int) {
    numChildren := input[idx]
    numMetadata := input[idx+1]
    idx += 2
    metadataValue := 0
    metadataValues := make([]int, 0)
    values := make([]int, 0)
    var value int
    for i := 0; i < numChildren; i++ {
        idx, value = buildTreeRecursive(input, idx)
        values = append(values, value)
    }

    for i := 0; i < numMetadata; i++ {
        metadataValues = append(metadataValues, input[idx])
        idx++
    }

    if numChildren == 0 {
        for i := 0; i < len(metadataValues); i++ {
            metadataValue += metadataValues[i]
        }
    } else {
        for i := 0; i < len(metadataValues); i++ {
            if metadataValues[i] != 0 && metadataValues[i] <= numChildren {
                metadataValue += values[metadataValues[i]-1]
            }
        }
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

    _, metadataValue := buildTreeRecursive(input, 0)
    fmt.Println(metadataValue)
}
