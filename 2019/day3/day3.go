package main

import (
	"bufio"
	"fmt"
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

func intersectLineWithChain(line string, startLine []int, chain []string, startchain []int) (intersect bool, intersections [][]int) {
	intersect = false
	endLine := advance(line, startLine)
	for _, chainLine := range chain {
		endChain := advance(chainLine, startchain)
		t := float64((startLine[0]-startchain[0])*(startchain[1]-endChain[1])-
			(endLine[1]-startchain[1])*(startchain[0]-endChain[0])) /
			float64((startLine[0]-endLine[0])*(startchain[1]-endChain[1])-
				(startLine[1]-endLine[1])*(startchain[0]-endChain[0]))
		if t <= 1 && t >= 0 {
			intersectX := startLine[0] + int(t*float64(endLine[0]-startLine[0]))
			intersectY := startLine[1] + int(t*float64(endLine[1]-startLine[1]))
			intersections = append(intersections, []int{intersectX, intersectY})
		}
		copy(startchain, endChain)
	}
	fmt.Println(intersections)
	return
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

	fmt.Println(firstWire)
	fmt.Println(secondWire)

	start := []int{0, 0}

	intersectLineWithChain("R100", start, []string{"D100", "R50", "U100"}, []int{10, 10})
}
