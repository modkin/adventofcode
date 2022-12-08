package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
)

func look(trees [][]int, incx int, incy int) (freeView bool, treesVisible int) {

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
	score := 0
	fmt.Println(counter)
	for y := 1; y < len(trees)-1; y++ {
		//outer:

		for x := 1; x < len(trees[0])-1; x++ {
			currentScore := 1
			currentHeight := trees[y][x]
			factor := 0
			for i := x - 1; i >= 0; i-- {
				factor++
				if trees[y][i] >= currentHeight {
					break
				}

				//if i == 0 {
				//	counter++
				//	continue outer
				//}
			}
			currentScore *= factor
			factor = 0
			for i := x + 1; i <= len(trees[0])-1; i++ {
				factor++
				if trees[y][i] >= currentHeight {
					break
				}

				//if i == len(trees[0])-1 {
				//	counter++
				//	continue outer
				//}
			}
			currentScore *= factor
			factor = 0
			for i := y - 1; i >= 0; i-- {
				factor++
				if trees[i][x] >= currentHeight {
					break
				}

				//if i == 0 {
				//	counter++
				//	continue outer
				//}
			}
			currentScore *= factor
			factor = 0
			for i := y + 1; i <= len(trees[0])-1; i++ {
				factor++
				if trees[i][x] >= currentHeight {
					break
				}

				//if i == len(trees)-1 {
				//	counter++
				//	continue outer
				//}
			}
			currentScore *= factor
			if currentScore > score {
				score = currentScore
				fmt.Println(score, x, y)
			}
		}

	}

	fmt.Println("Day 7.1:", counter)

}
