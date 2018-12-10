package main

import (
    "fmt"
    "regexp"
    "bufio"
    "strconv"
    "os"
    "math"
)

type Vector struct {
    px int
    py int
    vx int
    vy int
}

func draw(vector []*Vector) {
    seconds := 100000
    for seconds >= 0 {
        maxX, minX, maxY, minY := getBounds(vector)
        xi := minX
        yi := minY
        if minX + 100 >= maxX && minY + 100 >= maxY {
            for xi <= maxX {
                for yi <= maxY {
                    found := false
                    for _, vc := range vector {
                        if vc.px == xi && vc.py == yi {
                            fmt.Printf("#")
                            found = true
                            break;
                        }
                    }
                    if !found {
                        fmt.Printf(".")
                    }
                    yi++
                }
                fmt.Println()
                xi++
                yi = minY
            }
        }
        for _, vc := range vector {
            vc.px += vc.vx
            vc.py += vc.vy
        }
        seconds--
    }
}

func getBounds(vector []*Vector) (int, int, int, int) {
    maxX := 0
    minX := 999999
    maxY := 0
    minY := 999999
    for _, vc:= range vector {
        maxX = getMax(maxX, vc.px) 
        minX = getMin(minX, vc.px)
        maxY = getMax(maxY, vc.py) 
        minY = getMin(minY, vc.py)
    }
    return maxX, minX, maxY, minY
}

func getMax(num1, num2 int) int {
    return int(math.Max(float64(num1), float64(num2)))
}

func getMin(num1, num2 int) int {
    return int(math.Min(float64(num1), float64(num2)))
}


func main() {
    f, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(f)
    r := regexp.MustCompile("-*[0-9]+")
    vector := make([]*Vector, 0)
    for scanner.Scan() {
        matches := r.FindAllString(scanner.Text(), -1)
        px, _ := strconv.Atoi(matches[0])
        py, _ := strconv.Atoi(matches[1])
        vx, _ := strconv.Atoi(matches[2])
        vy, _ := strconv.Atoi(matches[3])
        vector = append(vector, &Vector{px, py, vx, vy})
    }
    draw(vector)
}
