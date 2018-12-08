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

func isInfinite(x, y, maxX, maxY int) bool {
    return x == 0 || y == 0 || x == maxX || y == maxY
}

func isNonMinEquidistant(minDistance int, distsFromPoint []int) bool {
        minOccurrences := 0
        for _, dist := range distsFromPoint {
            if dist == minDistance {
                minOccurrences++
            }
        }
        return minOccurrences < 2
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

    m := make(map[int]int)
    for i := 0; i <= maxX; i++ {
        for j:= 0; j <= maxY; j++ {
            closestPointIndex := -1
            minDistFromPoint := math.MaxInt32
            distsFromPoint := make([]int,0)
            for k := 0; k < len(coords); k++ {
                distFromPoint := manhattanDistance(i,j, coords[k].x, coords[k].y)
                distsFromPoint = append(distsFromPoint, distFromPoint)

                if distFromPoint < minDistFromPoint {
                    minDistFromPoint = distFromPoint
                    closestPointIndex = k
                }
            }


            if !isNonMinEquidistant(minDistFromPoint, distsFromPoint) {
                continue
            }

            if isInfinite(i, j, maxX, maxY) {
                m[closestPointIndex] = -math.MaxInt32
            } else {
                m[closestPointIndex]++
            }
        }
    }

    maxArea := 0
    for i := 0; i < len(coords); i++ {
        maxArea = int(math.Max(float64(maxArea), float64(m[i])))
    }

    fmt.Println(maxArea)
}
