package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func turnLeft(vec [2]int) (ret [2]int) {
	ret[0] = vec[1]
	ret[1] = -1 * vec[0]
	return
}

func turnRight(vec [2]int) (ret [2]int) {
	ret[0] = -1 * vec[1]
	ret[1] = vec[0]
	return
}

func colMinMax(grid map[[2]int]string, x int) (min, max int) {
	outside := true
	for y := 0; y < math.MaxInt; y++ {
		if outside {
			if _, ok := grid[[2]int{x, y}]; ok {
				min = y
				outside = false
			}
		} else {
			if _, ok := grid[[2]int{x, y}]; !ok {
				max = y - 1
				return
			}
		}
	}
	fmt.Println("ERROR")
	return -math.MaxInt, math.MaxInt
}

func rowMinMax(grid map[[2]int]string, y int) (min, max int) {
	outside := true
	for x := 0; x < math.MaxInt; x++ {
		if outside {
			if _, ok := grid[[2]int{x, y}]; ok {
				min = x
				outside = false
			}
		} else {
			if _, ok := grid[[2]int{x, y}]; !ok {
				max = x - 1
				return
			}
		}
	}
	fmt.Println("ERROR")
	return -math.MaxInt, math.MaxInt
}

func add(first [2]int, second [2]int) [2]int {
	return [2]int{first[0] + second[0], first[1] + second[1]}
}

func move(grid map[[2]int]string, start [2]int, dir [2]int, steps int) (pos [2]int) {
	pos = start
	for i := 0; i < steps; i++ {
		if tile, ok := grid[add(pos, dir)]; ok {
			if tile == "#" {
				return
			} else {
				pos = add(pos, dir)
			}
		} else {
			if dir[0] != 0 {
				min, max := rowMinMax(grid, pos[1])
				if dir[0] == 1 {
					if tar := [2]int{min, pos[1]}; grid[tar] == "#" {
						return
					} else {
						pos = tar
					}
				} else if dir[0] == -1 {
					if tar := [2]int{max, pos[1]}; grid[tar] == "#" {
						return
					} else {
						pos = tar
					}
				}
			} else {
				min, max := colMinMax(grid, pos[0])
				if dir[1] == 1 {
					if tar := [2]int{pos[0], min}; grid[tar] == "#" {
						return
					} else {
						pos = tar
					}
				} else if dir[1] == -1 {
					if tar := [2]int{pos[0], max}; grid[tar] == "#" {
						return
					} else {
						pos = tar
					}
				}
			}
		}
	}
	return pos

}

func main() {
	file, err := os.Open("2022/day22/input")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	grid := make(map[[2]int]string)
	path := make([]string, 0)

	yPos := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			scanner.Scan()
			line := scanner.Text()
			numreg := regexp.MustCompile(`\d*`)
			dirreg := regexp.MustCompile(`R|L`)
			numbers := numreg.FindAllString(line, -1)
			dirs := dirreg.FindAllString(line, -1)
			for i, dir := range dirs {
				path = append(path, numbers[i])
				path = append(path, dir)
			}
			path = append(path, numbers[len(numbers)-1])
		} else {

			line := strings.Split(scanner.Text(), "")
			for xPos, s := range line {
				if s != " " {
					grid[[2]int{xPos, yPos}] = s
				}
			}
			yPos++
		}
	}
	fmt.Println(path)
	utils.Print2DStringsGrid(grid)
	startx, _ := rowMinMax(grid, 0)
	pos := [2]int{startx, 0}
	dir := [2]int{1, 0}
	for i, s := range path {
		if i%2 == 0 {
			steps := utils.ToInt(s)
			pos = move(grid, pos, dir, steps)
		} else {
			if s == "L" {
				dir = turnLeft(dir)
			} else {
				dir = turnRight(dir)
			}
			grid[pos] = "*"
			//utils.Print2DStringsGrid(grid)
		}
	}
	row := pos[1] + 1
	col := pos[0] + 1
	password := row*1000 + col*4
	if dir == [2]int{1, 0} {
		password += 0
	} else if dir == [2]int{0, 1} {
		password += 1
	} else if dir == [2]int{-1, 0} {
		password += 2
	} else if dir == [2]int{0, -1} {
		password += 3
	}
	fmt.Println(password)
}
