package main

import (
	"adventofcode/2018/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

///   ^
///   |
///   y
///   |
///   0---x---->
/// [x][y]

func advance(line string, start []int) []int {
	end := make([]int, len(start))
	copy(end, start)
	step, _ := strconv.Atoi(line[1:])
	if line[0] == 'R' {
		end[0] += step
	} else if line[0] == 'L' {
		end[0] -= step
	} else if line[0] == 'U' {
		end[1] += step
	} else if line[0] == 'D' {
		end[1] -= step
	}
	return end
}

func intersectLineWithChain(line string, startLine []int, chain []string, startChain []int) [][2]int {
	var intersections [][2]int
	endLine := advance(line, startLine)
	for _, chainLine := range chain {
		endChain := advance(chainLine, startChain)
		t := float64((startLine[0]-startChain[0])*(startChain[1]-endChain[1])-
			(startLine[1]-startChain[1])*(startChain[0]-endChain[0])) /
			float64((startLine[0]-endLine[0])*(startChain[1]-endChain[1])-
				(startLine[1]-endLine[1])*(startChain[0]-endChain[0]))
		u := -float64((startLine[0]-endLine[0])*(startLine[1]-startChain[1])-
			(startLine[1]-endLine[1])*(endLine[0]-startChain[0])) /
			float64((startLine[0]-endLine[0])*(startChain[1]-endChain[1])-
				(startLine[1]-endLine[1])*(startChain[0]-endChain[0]))
		if t <= 1 && t >= 0 && u <= 1 && u >= 0 {
			intersectX := startLine[0] + int(math.Round(t*float64(endLine[0]-startLine[0])))
			intersectY := startLine[1] + int(math.Round(t*float64(endLine[1]-startLine[1])))
			intersections = append(intersections, [2]int{intersectX, intersectY})
		}
		copy(startChain, endChain)
	}
	//fmt.Println(intersections)
	return intersections
}

func shortestDistance(intersections [][2]int) int {
	min := math.MaxInt32
	for _, elem := range intersections {
		distance := utils.IntAbs(elem[0]) + utils.IntAbs(elem[1])
		if distance < min && distance != 0 {
			min = distance
		}
	}
	return min
}

func main() {
	file, err := os.Open("2019/day3/input")
	if err != nil {
		panic(err)
	}

	//var firstWire []string
	//var secondWire []string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstWire := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	secondWire := strings.Split(scanner.Text(), ",")

	//firstWire = strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72",",")
	//secondWire = strings.Split("U62,R66,U55,R34,D71,R55,D58,R83",",")
	//firstWire = strings.Split("R8,U5,L5,D3", ",")
	//secondWire = strings.Split("U7,R6,D4,L4", ",")

	//fmt.Println(firstWire)
	//fmt.Println(secondWire)

	start := []int{0, 0}

	var intersections [][2]int
	for _, line := range firstWire {
		intersections = append(intersections, intersectLineWithChain(line, start, secondWire, []int{0, 0})...)
		start = advance(line, start)
	}
	//fmt.Println(intersections)
	fmt.Println("Task 3.1: ", shortestDistance(intersections))
}
