package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type plotType struct {
	area      int
	allArea   map[[2]int]bool
	perimeter map[[2]float64]bool
}

func clacSides(plot plotType) int {
	//checked := make(map[[2]float64]bool)

	input := plot.perimeter
	counter := 0
	for len(input) > 0 {
		for pos1, _ := range input {
			toDelete := [][2]float64{}
			toDelete = append(toDelete, pos1)

			if math.Abs(math.Mod(pos1[0], 1)) == 0.5 {

				xPos := pos1[1] - 1
				for {
					if plot.allArea[[2]int{int(pos1[0] - 0.5), int(pos1[1])}] == plot.allArea[[2]int{int(pos1[0] - 0.5), int(xPos)}] ||
						plot.allArea[[2]int{int(pos1[0] + 0.5), int(pos1[1])}] == plot.allArea[[2]int{int(pos1[0] + 0.5), int(xPos)}] {
						if _, ok := input[[2]float64{pos1[0], xPos}]; ok {
							toDelete = append(toDelete, [2]float64{pos1[0], xPos})
							xPos--
						} else {
							break
						}
					} else {
						break
					}
				}
				xPos = pos1[1] + 1
				for {
					if plot.allArea[[2]int{int(pos1[0] - 0.5), int(pos1[1])}] == plot.allArea[[2]int{int(pos1[0] - 0.5), int(xPos)}] ||
						plot.allArea[[2]int{int(pos1[0] + 0.5), int(pos1[1])}] == plot.allArea[[2]int{int(pos1[0] + 0.5), int(xPos)}] {
						if _, ok := input[[2]float64{pos1[0], xPos}]; ok {
							toDelete = append(toDelete, [2]float64{pos1[0], xPos})
							xPos++
						} else {
							break
						}
					} else {
						break
					}
				}
			} else {
				yPos := pos1[0] - 1
				for {
					if plot.allArea[[2]int{int(pos1[0]), int(pos1[1] - 0.5)}] == plot.allArea[[2]int{int(yPos), int(pos1[1] - 0.5)}] ||
						plot.allArea[[2]int{int(pos1[0]), int(pos1[1] + 0.5)}] == plot.allArea[[2]int{int(yPos), int(pos1[1] + 0.5)}] {
						if _, ok := input[[2]float64{yPos, pos1[1]}]; ok {
							toDelete = append(toDelete, [2]float64{yPos, pos1[1]})
							yPos--
						} else {
							break
						}
					} else {
						break
					}
				}
				yPos = pos1[0] + 1
				for {
					if plot.allArea[[2]int{int(pos1[0]), int(pos1[1] - 0.5)}] == plot.allArea[[2]int{int(yPos), int(pos1[1] - 0.5)}] ||
						plot.allArea[[2]int{int(pos1[0]), int(pos1[1] + 0.5)}] == plot.allArea[[2]int{int(yPos), int(pos1[1] + 0.5)}] {
						if _, ok := input[[2]float64{yPos, pos1[1]}]; ok {
							toDelete = append(toDelete, [2]float64{yPos, pos1[1]})
							yPos++
						} else {
							break
						}
					} else {
						break
					}
				}
			}

			for _, p := range toDelete {
				delete(input, p)
			}
			counter++
			break
		}
	}
	return counter
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

			//utils.Print2DStringGrid(plantArea)

			typeCounter[plant] = typeCounter[plant] + 1
		}
	}

	for y, line := range input {
		for x, plant := range line {
			if _, ok := allPlots[plant]; !ok {
				allPlots[plant] = plotType{allArea: make(map[[2]int]bool), perimeter: make(map[[2]float64]bool)}
			}
			plot := allPlots[plant]
			plot.area = plot.area + 1
			plot.allArea[[2]int{y, x}] = true
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
	sum2 := 0
	for t, b := range allPlots {
		if b.area != len(b.allArea) {
			fmt.Println("ERROR")
		}
		sum += b.area * len(b.perimeter)
		sides := clacSides(b)
		sum2 += b.area * sides
		fmt.Println(t, sides, b.area*len(b.perimeter), b.area, len(b.perimeter))
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}
