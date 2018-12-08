package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "math"
)

type Coord struct {
    x int
    y int
}

func manhattanDistance(x1, y1, x2, y2 int) int {
    return int(math.Abs(float64(x1) - float64(x2)) + math.Abs(float64(y1) - float64(y2)))
}

func main() {
    f, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(f)
    r := regexp.MustCompile("[0-9]+")
    coords := make([]Coord, 0)
    maxX := 0
    maxY := 0
    for scanner.Scan() {
        match := r.FindAllString(scanner.Text(), -1)
        x, _ := strconv.Atoi(match[0])
        y, _ := strconv.Atoi(match[1])
        coords = append(coords, Coord{x, y})
        maxX = int(math.Max(float64(maxX), float64(x)))
        maxY = int(math.Max(float64(maxY), float64(y)))
    }

    m := make([][]int, maxX+1)
    for i := range m {
        m[i] = make([]int, maxY+1)
    }

    for i := 0; i <= maxX; i++ {
        for j:= 0; j <= maxY; j++ {
            for k := 0; k < len(coords); k++ {
                distance := manhattanDistance(i,j, coords[k].x, coords[k].y)
                m[i][j] += distance
            }
        }
    }
    area := 0
    for i := 0; i <= maxX; i++ {
        for j:= 0; j <= maxY; j++ {
            if m[i][j] < 10000 {
                area++
            }
        }
    }
    fmt.Println(area)
}
