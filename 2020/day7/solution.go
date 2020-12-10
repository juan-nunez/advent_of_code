package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BagProperty struct {
	name string
	num  int
}

var graph = make(map[string][]BagProperty, 0)
var visited = make(map[string]bool, 0)

func buildGraph() map[string][]BagProperty {
	f, _ := os.Open("input.txt")
	defer f.Close()

	re := regexp.MustCompile(`([0-9]+) ([a-z]+ [a-z]+ bag)`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		node := strings.Join(split[0:3], " ")
		node = node[:len(node)-1]
		visited[node] = false
		graph[node] = []BagProperty{}
		parts := re.FindAllStringSubmatch(scanner.Text(), -1)
		for i := 0; i < len(parts); i++ {
			numBags, _ := strconv.Atoi(parts[i][1])
			bagName := parts[i][2]
			graph[node] = append(graph[node], BagProperty{bagName, numBags})
		}
	}

	return graph
}

func dfs(node string) bool {
	if node == "shiny gold bag" {
		return true
	}

	found := false
	visited[node] = true
	edges := graph[node]

	for i := 0; i < len(edges); i++ {
		edgeNode := edges[i].name
		if !visited[edgeNode] {
			found = dfs(edgeNode)
			if found {
				return true
			}
		}
	}

	return false
}

func doPartOne() int {
	buildGraph()
	ans := 0
	for k, _ := range graph {
		found := dfs(k)
		if found {
			ans++
		}
		for key, _ := range visited {
			visited[key] = false
		}
	}

	return ans - 1
}

func countBags(node string) int {
	visited[node] = true
	edges := graph[node]

	numBags := 0
	for i := 0; i < len(edges); i++ {
		edgeNode := edges[i].name
		numBags += edges[i].num + edges[i].num*countBags(edgeNode)
	}

	return numBags
}

func doPartTwo() int {
	buildGraph()
	ans := countBags("shiny gold bag")
	return ans
}

func main() {
	fmt.Println(doPartOne())

	// Reset graph state
	graph = make(map[string][]BagProperty, 0)
	visited = make(map[string]bool, 0)

	fmt.Println(doPartTwo())
}
