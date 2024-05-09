package day25

import (
	_ "embed"
	"github.com/ictrobot/aoc-go/internal/util/parse"
	"math/rand"
	"strings"
)

//go:embed example
var Example string

type Day25 struct {
	edges [][]int
}

func (d *Day25) Parse(input string) {
	lines := parse.Grid(parse.Lines, parse.Whitespace, strings.ReplaceAll(input, ":", ""))

	names := make(map[string]int)
	for _, line := range lines {
		for _, name := range line {
			if _, exists := names[name]; !exists {
				names[name] = len(names)
			}
		}
	}

	d.edges = make([][]int, len(names))
	for _, line := range lines {
		lhs := names[line[0]]
		for _, name := range line[1:] {
			rhs := names[name]
			d.edges[lhs] = append(d.edges[lhs], rhs)
			d.edges[rhs] = append(d.edges[rhs], lhs)
		}
	}
}

func (d *Day25) ParseExample() {
	d.Parse(Example)
}

func (d *Day25) Part1() any {
	for {
		removedEdges := d.edgesToRemove()
		count1, count2, ok := d.checkSplit(removedEdges)
		if ok {
			return count1 * count2
		}
	}
}

func (d *Day25) Part2() any {
	// no part 2
	return "🎄"
}

func (d *Day25) edgesToRemove() (removedEdges [3][2]int) {
	edgeCount := make(map[[2]int]int)
	visited := make([]bool, len(d.edges))

	var dfs func(int, int)
	dfs = func(v, removed int) {
		visited[v] = true

		// rand.Perm is needed to avoid always visiting the same edge first
		// from a given vertex which leads to almost always visiting nodes
		// in the same order
	neighbours:
		for _, p := range rand.Perm(len(d.edges[v])) {
			neighbor := d.edges[v][p]
			if visited[neighbor] {
				continue
			}

			edge := [2]int{min(v, neighbor), max(v, neighbor)}
			for _, e := range removedEdges[:removed] {
				if e == edge {
					continue neighbours
				}
			}
			edgeCount[edge]++

			dfs(neighbor, removed)
		}
	}

	for i := 0; i < 3; i++ {
		clear(edgeCount)

		for source := range d.edges {
			clear(visited)
			dfs(source, i)
		}

		var maxCount int
		for e, c := range edgeCount {
			if c > maxCount {
				maxCount = c
				removedEdges[i] = e
			}
		}
	}

	return
}

func (d *Day25) checkSplit(removedEdges [3][2]int) (int, int, bool) {
	visited := make([]bool, len(d.edges))

	var countReachable func(int) int
	countReachable = func(v int) int {
		visited[v] = true
		count := 1

		for _, neighbor := range d.edges[v] {
			if visited[neighbor] {
				continue
			}

			edge := [2]int{min(v, neighbor), max(v, neighbor)}
			if removedEdges[0] == edge || removedEdges[1] == edge || removedEdges[2] == edge {
				continue
			}

			count += countReachable(neighbor)
		}

		return count
	}

	count1 := countReachable(0)

	unvisited := -1
	for i, value := range visited {
		if !value {
			unvisited = i
			break
		}
	}
	if unvisited == -1 {
		return 0, 0, false
	}

	count2 := countReachable(unvisited)

	if count1+count2 < len(d.edges) {
		return 0, 0, false
	}
	return count1, count2, true
}
