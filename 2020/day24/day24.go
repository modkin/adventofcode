package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"strings"
)

func printTiles(paintMap map[[2]int]bool) {
	maxX, maxY := math.MinInt32, math.MinInt32
	for pos := range paintMap {
		if pos[0] > maxX {
			maxX = pos[0]
		}
		if pos[1] > maxY {
			maxY = pos[1]
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if value, ok := paintMap[[2]int{x, y}]; ok {
				if value {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
}

func move(start [2]int, direction string) [2]int {
	if direction == "e" {
		return [2]int{start[0] + 1, start[1]}
	} else if direction == "w" {
		return [2]int{start[0] - 1, start[1]}
	} else if direction == "nw" {
		return [2]int{start[0], start[1] + 1}
	} else if direction == "ne" {
		return [2]int{start[0] + 1, start[1] + 1}
	} else if direction == "sw" {
		return [2]int{start[0] - 1, start[1] - 1}
	} else if direction == "se" {
		return [2]int{start[0], start[1] - 1}
	}
	fmt.Println("ERROR")
	return [2]int{math.MaxInt64, math.MaxInt64}
}

func countNeighbors(pos [2]int, tiles map[[2]int]bool) int {
	allDirections := []string{"e", "w", "ne", "nw", "se", "sw"}
	neighbors := 0
	for _, dir := range allDirections {
		nbr := move(pos, dir)
		if tiles[nbr] {
			neighbors++
		}
	}
	return neighbors
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day24/input"))
	tiles := make(map[[2]int]bool)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		pos := [2]int{0, 0}
		skipNext := false
		for i, elem := range line {
			if skipNext {
				skipNext = false
				continue
			}
			if elem == "e" || elem == "w" {
				pos = move(pos, elem)
			} else {
				skipNext = true
				pos = move(pos, elem+line[i+1])
			}
		}
		if tiles[pos] {
			delete(tiles, pos)
		} else {
			tiles[pos] = true
		}
	}
	printTiles(tiles)
	countTiles := func() int {
		count := 0
		for _, value := range tiles {
			if value {
				count++
			}
		}
		return count
	}

	fmt.Println(countTiles())

	for i := 0; i < 100; i++ {
		newtiles := make(map[[2]int]bool)
		for position, _ := range tiles {
			// check if white has two neighbors
			allDirections := []string{"e", "w", "ne", "nw", "se", "sw"}
			for _, dir := range allDirections {
				nbr := move(position, dir)
				if !tiles[nbr] {
					if countNeighbors(nbr, tiles) == 2 {
						newtiles[nbr] = true
					}
				}
			}
			// check black
			if nrNbr := countNeighbors(position, tiles); nrNbr == 1 || nrNbr == 2 {
				newtiles[position] = true
			}
		}
		tiles = newtiles
		fmt.Println(countTiles())
	}
}
