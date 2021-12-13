package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2021/day13/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	paper := make(map[[2]int]bool)
	folds := make([]string, 0)
	first := true
	ymax, xmax := 0, 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := strings.Split(scanner.Text(), ",")
		x, y := utils.ToInt(line[0]), utils.ToInt(line[1])
		newPoint := [2]int{x, y}
		paper[newPoint] = true
		if x > xmax {
			xmax = x
		}
		if y > ymax {
			ymax = y
		}
	}
	//utils.Print2DStringGrid(paper)
	for scanner.Scan() {
		folds = append(folds, scanner.Text())
	}

	for _, fold := range folds {
		coord := utils.ToInt(strings.Split(strings.Split(fold, " ")[2], "=")[1])
		if strings.Contains(fold, "y") {
			foldBy := -2
			for y := coord + 1; y <= ymax; y++ {
				for x := 0; x <= xmax; x++ {
					if _, ok := paper[[2]int{x, y}]; ok {
						paper[[2]int{x, y + foldBy}] = true
						delete(paper, [2]int{x, y})
					}
				}
				foldBy -= 2
			}
			ymax = coord - 1
		} else {
			foldBy := -2

			for x := coord + 1; x <= xmax; x++ {
				for y := 0; y <= ymax; y++ {
					if _, ok := paper[[2]int{x, y}]; ok {
						paper[[2]int{x + foldBy, y}] = true
						delete(paper, [2]int{x, y})
					}
				}
				foldBy -= 2
			}
			xmax = coord - 1
		}
		if first {
			fmt.Println("Day 13.1:", len(paper))
			first = false
		}
		//utils.Print2DStringGrid(paper)
	}
	utils.Print2DStringGrid(paper)
}
