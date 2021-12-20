package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func toInt(input []string) int {
	number, err := strconv.ParseInt(strings.Join(input, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	return int(number)
}

func enhance(grid map[[2]int]string, algo []string) map[[2]int]string {
	output := make(map[[2]int]string)
	xMin, yMin, xMax, yMax := math.MaxInt, math.MaxInt, 0, 0
	for i := range grid {
		if i[0] < xMin {
			xMin = i[0]
		}
		if i[0] > xMax {
			xMax = i[0]
		}
		if i[1] < yMin {
			yMin = i[1]
		}
		if i[1] > yMax {
			yMax = i[1]
		}
	}

	for y := yMin - 10; y <= yMax+10; y++ {
		for x := xMin - 10; x <= xMax+10; x++ {
			enhancePos := make([]string, 0)
			for i := -1; i <= 1; i++ {
				for m := -1; m <= 1; m++ {
					pos := [2]int{x + m, y + i}
					if val, ok := grid[pos]; ok {
						if val == "#" {
							enhancePos = append(enhancePos, "1")
						} else {
							enhancePos = append(enhancePos, "0")
						}
					} else {
						enhancePos = append(enhancePos, "0")
					}
				}
			}
			intPos := toInt(enhancePos)
			output[[2]int{x, y}] = algo[intPos]
		}
	}
	return output
}

func main() {
	file, err := os.Open("2021/day20/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	scanner.Scan()

	algo := make([]string, 0)
	algo = strings.Split(scanner.Text(), "")
	scanner.Scan()
	image := make(map[[2]int]string)
	y := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for i, l := range line {
			newEntry := [2]int{i, y}
			image[newEntry] = l
		}
		y++
	}
	utils.Print2DStringsGrid(image)
	//fmt.Println(algo)
	image = enhance(image, algo)
	image = enhance(image, algo)
	utils.Print2DStringsGrid(image)
	counter := 0
	for key, i2 := range image {
		if key[0] > -5 && key[0] < 105 && key[1] > -5 && key[1] < 105 {
			if i2 == "#" {
				counter++
			}
		}
	}
	fmt.Println(counter)

}
