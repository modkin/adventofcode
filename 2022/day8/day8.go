package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
)

func look(trees [][]int, x int, y int, incx int, incy int) (freeView bool, treesVisible int) {
	treeHeight := trees[y][x]
	treesVisible = 1
	for {
		x += incx
		y += incy
		if trees[y][x] >= treeHeight {
			return false, treesVisible
		}
		if x == 0 || x == len(trees[0])-1 || y == 0 || y == len(trees)-1 {
			return true, treesVisible
		}
		treesVisible++
	}
}

func main() {

	file, err := os.Open("2022/day8/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	trees := make([][]int, 0)

	for scanner.Scan() {
		split := []rune(scanner.Text())
		tmp := make([]int, 0)
		for _, r := range split {
			tmp = append(tmp, utils.ToInt(string(r)))
		}
		trees = append(trees, tmp)
	}

	counter := len(trees) * 2
	counter += (len(trees[0]) * 2) - 4
	maxScenicScore := 0
	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees[0])-1; x++ {

			leftView, leftTrees := look(trees, x, y, -1, 0)
			rightView, rightTrees := look(trees, x, y, 1, 0)
			bottomView, bottomTrees := look(trees, x, y, 0, -1)
			topView, topTrees := look(trees, x, y, 0, 1)

			if leftView || rightView || bottomView || topView {
				counter++
			}
			currentScenicScore := leftTrees * rightTrees * bottomTrees * topTrees

			if currentScenicScore > maxScenicScore {
				maxScenicScore = currentScenicScore
			}
		}
	}

	fmt.Println("Day 8.1:", counter)
	fmt.Println("Day 8.2:", maxScenicScore)
}
