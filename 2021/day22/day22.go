package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2021/day22/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	cubes := make(map[[3]int]bool)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		onOff := strings.Split(line[0], " ")[0]
		line[0] = strings.Trim(strings.Trim(line[0], "on "), "off ")
		xArea := strings.Split(strings.Trim(line[0], "x="), "..")
		yArea := strings.Split(strings.Trim(line[1], "y="), "..")
		zArea := strings.Split(strings.Trim(line[2], "z="), "..")
		fmt.Println(onOff, xArea, yArea, zArea)
		for x := utils.ToInt(xArea[0]); x <= utils.ToInt(xArea[1]) && utils.IntAbs(x) <= 50; x++ {
			for y := utils.ToInt(yArea[0]); y <= utils.ToInt(yArea[1]) && utils.IntAbs(y) <= 50; y++ {
				for z := utils.ToInt(zArea[0]); z <= utils.ToInt(zArea[1]) && utils.IntAbs(z) <= 50; z++ {

					if onOff == "on" {
						cubes[[3]int{x, y, z}] = true
					} else {
						cubes[[3]int{x, y, z}] = false
					}

				}
			}
		}
	}
	counter := 0
	for key, val := range cubes {
		if utils.IntAbs(key[0]) <= 50 && utils.IntAbs(key[1]) <= 50 && utils.IntAbs(key[2]) <= 50 {
			if val == true {
				counter++
			}
		}
	}
	fmt.Println(counter)
}
