package main

import (
	"adventofcode/utils"
	"container/heap"
	"fmt"
	"strings"
)

func AddPoints(one, two [2]int) [2]int {
	return [2]int{one[0] + two[0], one[1] + two[1]}
}

func SubPoints(one, two [2]int) [2]int {
	return [2]int{two[0] - one[0], two[1] - one[1]}
}

type posType struct {
	pos  [2]int
	cost int
}

type posHeap []posType

func (h posHeap) Len() int           { return len(h) }
func (h posHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h posHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *posHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(posType))
}

func (h *posHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day18/input")
	test := false
	var xMax, yMax int
	var stop int
	if test {
		xMax = 7
		yMax = 7
		stop = 12

	} else {
		xMax = 71
		yMax = 71
		stop = 1024
	}

	memory := make(map[[2]int]bool)
	start := [2]int{0, 0}
	target := [2]int{xMax - 1, yMax - 1}
	for i, line := range lines {

		if i >= stop {
			break
		}
		split := strings.Split(line, ",")
		memory[[2]int{utils.ToInt(split[0]), utils.ToInt(split[1])}] = true

	}
	utils.Print2DStringGrid(memory)

	fmt.Println(start, target)

	h := &posHeap{}
	heap.Init(h)
	heap.Push(h, posType{pos: start, cost: 0})
	visited := make(map[[2]int]bool)
	visited[start] = true

	cost := make(map[[2]int]int)
	cost[start] = 0

	for {
		minP := heap.Pop(h).(posType)
		if minP.pos == target {
			fmt.Println(minP.cost)
			break
		}

		visited[minP.pos] = true
		for _, offset := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {

			nextPos := AddPoints(minP.pos, offset)
			if _, ok := visited[nextPos]; ok {
				continue
			}
			if _, ok := memory[nextPos]; ok {
				continue
			}
			if nextPos[0] < 0 || nextPos[0] >= xMax || nextPos[1] < 0 || nextPos[1] >= yMax {
				continue
			}

			nextCost := minP.cost + 1
			if val, ok := cost[nextPos]; !ok || val > nextCost {
				newP := posType{cost: nextCost, pos: nextPos}
				heap.Push(h, newP)
				cost[nextPos] = nextCost
			}
		}

		//fmt.Println(h)
	}
	utils.Print2DStringGrid(visited)
	fmt.Println(cost)
	fmt.Println(cost[target])

}
