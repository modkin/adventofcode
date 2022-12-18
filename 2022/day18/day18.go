package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func max(in [3]int) int {
	maxi := 0
	for _, i := range in {
		if i > maxi {
			maxi = i
		}
	}
	return maxi
}

func getNeighbors(in [3]int) [][3]int {
	out := make([][3]int, 0)
	for idx := 0; idx < 3; idx++ {
		for _, step := range []int{-1, 1} {
			nbr := in
			nbr[idx] += step
			out = append(out, nbr)
		}
	}
	return out
}

func isBoundary(in [3]int, max int) bool {
	for _, i := range in {
		if i == 0 || i == max {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("2022/day18/input")

	if err != nil {
		panic(err)
	}

	cubes := make(map[[3]int]bool)
	airAround := make([][3]int, 0)

	scanner := bufio.NewScanner(file)

	maximum := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		newCube := [3]int{utils.ToInt(line[0]), utils.ToInt(line[1]), utils.ToInt(line[2])}
		if tmp := max(newCube); tmp > maximum {
			maximum = tmp
		}
		cubes[newCube] = true
	}
	maximum = maximum + 1

	for c := range cubes {
		for idx := 0; idx < 3; idx++ {
			for _, step := range []int{-1, 1} {
				check := c
				check[idx] += step
				if _, ok := cubes[check]; !ok {
					airAround = append(airAround, check)
				}
			}
		}
	}

	fmt.Println("Day 18.1:", len(airAround))

	airOutside := make([][3]int, 0)
airLoop:
	for _, air := range airAround {
		curPos := [][3]int{air}
		visited := map[[3]int]bool{air: true}

		for len(curPos) > 0 {
			nextPos := make([][3]int, 0)
			for _, pos := range curPos {
				nbrs := getNeighbors(pos)
				for _, nbr := range nbrs {
					if isBoundary(nbr, maximum) {
						airOutside = append(airOutside, air)
						continue airLoop
					} else {
						if _, islava := cubes[nbr]; !islava {
							if _, ok := visited[nbr]; !ok {
								nextPos = append(nextPos, nbr)
								visited[nbr] = true
							}
						}
					}
				}
			}
			curPos = nextPos
		}
	}
	fmt.Println("Day 18.2:", len(airOutside))

}
