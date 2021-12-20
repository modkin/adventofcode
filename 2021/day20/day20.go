package main

import (
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

func enhance(grid map[[2]int]string, algo []string, index int) map[[2]int]string {
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

	for y := yMin - 1; y <= yMax+1; y++ {
		for x := xMin - 1; x <= xMax+1; x++ {
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
						if index%2 == 0 {
							enhancePos = append(enhancePos, "0")
						} else {
							enhancePos = append(enhancePos, "1")
						}
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
	image = enhance(image, algo, 0)
	image = enhance(image, algo, 1)
	counter := 0
	for _, i2 := range image {
		if i2 == "#" {
			counter++
		}
	}
	for i := 2; i < 50; i++ {
		image = enhance(image, algo, i)
	}
	fmt.Println("Day 20.1:", counter)
	counter = 0
	for _, i2 := range image {
		if i2 == "#" {
			counter++
		}
	}
	fmt.Println("Day 20.2:", counter)

}
