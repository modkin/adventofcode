package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func moveTails(rope [][2]int) {
	for i := 0; i < len(rope)-1; i++ {
		diff0 := rope[i][0] - rope[i+1][0]
		diff1 := rope[i][1] - rope[i+1][1]
		if utils.IntAbs(diff0) > 1 && utils.IntAbs(diff1) > 1 {
			if diff0 > 0 {
				rope[i+1][0] += diff0 - 1
			} else {
				rope[i+1][0] += diff0 + 1
			}
			if diff1 > 0 {
				rope[i+1][1] += diff1 - 1
			} else {
				rope[i+1][1] += diff1 + 1
			}
		} else if utils.IntAbs(diff0) > 1 {
			if diff0 > 0 {
				rope[i+1][0] += diff0 - 1
			} else {
				rope[i+1][0] += diff0 + 1
			}
			rope[i+1][1] = rope[i][1]
		} else if utils.IntAbs(diff1) > 1 {
			if diff1 > 0 {
				rope[i+1][1] += diff1 - 1
			} else {
				rope[i+1][1] += diff1 + 1
			}
			rope[i+1][0] = rope[i][0]
		}
	}
}

func main() {

	file, err := os.Open("2022/day9/input")
	if err != nil {
		panic(err)
	}

	//grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)

	steps := make([]string, 0)

	for scanner.Scan() {
		steps = append(steps, scanner.Text())

	}

	visited := make(map[[2]int]int)
	head := [2]int{0, 0}
	tail := [2]int{0, 0}
	visited[tail] = 1
	for _, step := range steps {
		dir := strings.Split(step, " ")[0]
		stepSize := utils.ToInt(strings.Split(step, " ")[1])
		for i := 0; i < stepSize; i++ {
			if dir == "R" {
				head[0] += 1
			} else if dir == "L" {
				head[0] -= 1
			} else if dir == "U" {
				head[1] += 1
			} else if dir == "D" {
				head[1] -= 1
			}
			diff0 := head[0] - tail[0]
			diff1 := head[1] - tail[1]
			if utils.IntAbs(diff0) > 1 {
				if diff0 > 0 {
					tail[0] += diff0 - 1
				} else {
					tail[0] += diff0 + 1
				}
				tail[1] = head[1]
			}
			if utils.IntAbs(diff1) > 1 {
				if diff1 > 0 {
					tail[1] += diff1 - 1
				} else {
					tail[1] += diff1 + 1
				}
				tail[0] = head[0]
			}
			visited[tail] = 1

		}
	}

	fmt.Println("Day 9.1:", len(visited))

	stepCounter := 0
	visited2 := make(map[[2]int]int)
	rope := make([][2]int, 10)
	for i := range rope {
		rope[i] = [2]int{0, 0}
	}
	for _, step := range steps {
		dir := strings.Split(step, " ")[0]
		stepSize := utils.ToInt(strings.Split(step, " ")[1])
		for i := 0; i < stepSize; i++ {
			if dir == "R" {
				rope[0][0] += 1
			} else if dir == "L" {
				rope[0][0] -= 1
			} else if dir == "U" {
				rope[0][1] += 1
			} else if dir == "D" {
				rope[0][1] -= 1
			}
			moveTails(rope)
			visited2[rope[9]] = 1
			stepCounter++

		}
	}
	fmt.Println("Day 9.2:", len(visited2))

}
