package main

import (
	"adventofcode/utils"
	"golang.org/x/exp/errors/fmt"
	"strconv"
	"strings"
)

type plotType struct {
	area      int
	perimeter map[[2]float64]bool
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day12/input")

	var input [][]string
	maxX := 0
	maxY := 0

	allPlots := make(map[string]plotType)
	for _, line := range lines {
		split := strings.Split(line, "")
		maxX = len(split)
		input = append(input, split)
		maxY++
	}

	allPlots[input[0][0]] = plotType{perimeter: make(map[[2]float64]bool)}

	typeCounter := make(map[string]int)

	allChecked := make(map[[2]int]bool)
	for yOut, line := range input {
		for xOut, plant := range line {
			if _, ok := allChecked[[2]int{yOut, xOut}]; ok {
				continue
			}

			plantArea := make(map[[2]int]bool)
			plantArea[[2]int{yOut, xOut}] = true
			added := true
			for added {
				added = false
				for pos, _ := range plantArea {
					x := pos[1]
					y := pos[0]
					if x != 0 && input[y][x-1] == plant {
						if _, ok := plantArea[[2]int{y, x - 1}]; !ok {
							plantArea[[2]int{y, x - 1}] = true
							added = true
						}
					}
					if x != maxX-1 && input[y][x+1] == plant {
						if _, ok := plantArea[[2]int{y, x + 1}]; !ok {
							plantArea[[2]int{y, x + 1}] = true
							added = true
						}
					}
					if y != 0 && input[y-1][x] == plant {
						if _, ok := plantArea[[2]int{y - 1, x}]; !ok {
							plantArea[[2]int{y - 1, x}] = true
							added = true
						}
					}
					if y != maxY-1 && input[y+1][x] == plant {
						if _, ok := plantArea[[2]int{y + 1, x}]; !ok {
							plantArea[[2]int{y + 1, x}] = true
							added = true
						}
					}
				}
			}
			for ints, _ := range plantArea {
				input[ints[0]][ints[1]] = input[ints[0]][ints[1]] + strconv.Itoa(typeCounter[plant])
				allChecked[ints] = true
			}
			fmt.Println(plant)
			utils.Print2DStringGrid(plantArea)
			fmt.Println(input)
			typeCounter[plant] = typeCounter[plant] + 1
		}
	}

	for y, line := range input {
		for x, plant := range line {
			if _, ok := allPlots[plant]; !ok {
				allPlots[plant] = plotType{perimeter: make(map[[2]float64]bool)}
			}
			plot := allPlots[plant]
			plot.area = plot.area + 1
			if x == 0 || input[y][x-1] != plant {
				plot.perimeter[[2]float64{float64(y), float64(x) - 0.5}] = true
			}
			if x == maxX-1 || input[y][x+1] != plant {
				plot.perimeter[[2]float64{float64(y), float64(x) + 0.5}] = true
			}
			if y == 0 || input[y-1][x] != plant {
				plot.perimeter[[2]float64{float64(y) - 0.5, float64(x)}] = true
			}
			if y == maxY-1 || input[y+1][x] != plant {
				plot.perimeter[[2]float64{float64(y) + 0.5, float64(x)}] = true
			}
			allPlots[plant] = plot
		}
	}

	sum := 0
	for t, b := range allPlots {
		fmt.Println(t, b.area*len(b.perimeter), b.area, len(b.perimeter), b.perimeter)
		sum += b.area * len(b.perimeter)
	}
	fmt.Println(sum)
}
