package main

import (
	"adventofcode/utils"
	"container/heap"
	"fmt"
	"math"
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

func Print2DPath(grid map[[2]int]int) {
	xMax, yMax := 0, 0
	for i := range grid {
		if i[0] > xMax {
			xMax = i[0]
		}
		if i[1] > yMax {
			yMax = i[1]
		}
	}
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				if val == -1 {
					fmt.Print("#")
				} else if val == -2 {
					fmt.Print("*")
				} else {
					fmt.Print(val % 10)
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {

	racetrack := make(map[[2]int]string)
	var start, end [2]int

	lines := utils.ReadFileIntoLines("2024/day20/input")
	for y, line := range lines {
		split := strings.Split(line, "")
		for x, s := range split {
			if s == "S" {
				racetrack[[2]int{x, y}] = "."
				start = [2]int{x, y}
			} else if s == "E" {
				racetrack[[2]int{x, y}] = "."
				end = [2]int{x, y}
			} else {
				racetrack[[2]int{x, y}] = s
			}
		}

	}

	cost := make(map[[2]int]int)
	h := &posHeap{}
	heap.Init(h)
	heap.Push(h, posType{pos: start, cost: 0})
	visited := make(map[[2]int]bool)
	cost[start] = 0

	try := func(maxTime int, cheatSec int, cheatDir [2]int) (bool, int) {
		//timer := 0
		for {
			//if timer > maxTime {
			//	return false, -1
			//}
			if h.Len() == 0 {
				return false, -1
			}
			minP := heap.Pop(h).(posType)
			if _, ok := visited[minP.pos]; ok {
				continue
			} else {
				visited[minP.pos] = true
			}
			if minP.pos == end {
				return true, minP.cost
			}
			if minP.pos == [2]int{10, 7} {
				fmt.Println("j")
			}
			if minP.cost > maxTime {
				return false, -1
			}
			if minP.cost == cheatSec+1 {
				if val, ok := racetrack[minP.pos]; !ok || val == "#" {
					continue
				}
			}

			for _, offset := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {

				nextPos := AddPoints(minP.pos, offset)
				nextCost := minP.cost + 1
				if nextCost == cheatSec && offset == cheatDir { //|| nextCost == cheatSec+1 {
					//if _, ok := racetrack[nextPos]; !ok {
					//	continue
					//}
					fmt.Println(cheatSec, offset)
				} else {
					if val, ok := racetrack[nextPos]; !ok || val == "#" {
						continue
					}
				}
				if val, ok := cost[nextPos]; !ok || val > nextCost {
					newP := posType{cost: nextCost, pos: nextPos}
					heap.Push(h, newP)
					cost[nextPos] = nextCost
				}
			}
			//timer++
			//fmt.Println(h)
		}
	}

	found, baseTime := try(math.MaxInt, math.MaxInt, [2]int{0, 0})

	counter := 0

	CheatCounter := make(map[int]int)

	fmt.Println(baseTime)

	for t := 0; t <= baseTime; t++ {
		for _, cheatDir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			cost = make(map[[2]int]int)
			cost[start] = 0
			visited = make(map[[2]int]bool)
			for h.Len() > 0 {
				h.Pop()
			}
			h.Push(posType{pos: start, cost: 0})

			found2, time := try(baseTime, t, cheatDir)

			for ints, s := range racetrack {
				if _, ok := cost[ints]; !ok {
					if s == "#" {
						cost[ints] = -1
					}
				} else {
					if s == "#" {
						cost[ints] = -2
					}
				}
			}
			//fmt.Println("--------------", t, time)
			//Print2DPath(cost)
			//fmt.Println("--------------")
			if found2 && time <= baseTime-100 {
				counter++
			}
			if found2 {
				CheatCounter[baseTime-time]++
			}
		}
	}

	fmt.Println("GO")
	for i, i2 := range CheatCounter {
		fmt.Println(i2, i)
	}
	fmt.Println("Day 20.1:", found, counter)
}
