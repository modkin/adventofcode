package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

type cube struct {
	onOff string
	xArea [2]int
	yArea [2]int
	zArea [2]int
}

func toIntArea(area []string) (output [2]int) {
	output[0] = utils.ToInt(area[0])
	output[1] = utils.ToInt(area[1])
	return
}

func inLine(val int, area [2]int) bool {
	if val >= area[0] && val <= area[1] {
		return true
	} else {
		return false
	}
}

func inArea(area1 [2]int, area [2]int) bool {
	if area1[0] >= area[0] && area1[1] <= area[1] {
		return true
	} else {
		return false
	}
}

func findStatus(x int, y int, z int, cubes []cube) bool {
	for i := len(cubes) - 1; i >= 0; i-- {
		if inLine(x, cubes[i].xArea) && inLine(y, cubes[i].yArea) && inLine(z, cubes[i].zArea) {
			if cubes[i].onOff == "on" {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func getLineIntersect(first [2]int, second [2]int) (output [][2]int) {
	if inArea(first, second) {
		output = append(output, second)
	} else if inArea(second, first) {
		output = append(output, [2]int{first[0], second[0] - 1})
		output = append(output, [2]int{second[0], second[1]})
		output = append(output, [2]int{second[1] + 1, first[1]})
	} else {
		if inLine(first[0], second) {
			output = append(output, [2]int{second[0], second[1]})
			output = append(output, [2]int{second[1] + 1, first[1]})
		} else if inLine(first[1], second) {
			output = append(output, [2]int{first[0], second[0] - 1})
			output = append(output, [2]int{second[0], second[1]})
		}
	}
	return
}

//assuming two cubes intersect
func getIntersectCube(first cube, second cube) (output []cube) {
	xSections := getLineIntersect(first.xArea, second.xArea)
	ySections := getLineIntersect(first.yArea, second.yArea)
	zSections := getLineIntersect(first.zArea, second.zArea)
	for _, xs := range xSections {
		for _, ys := range ySections {
			for _, zs := range zSections {
				output = append(output, cube{"off", xs, ys, zs})
			}
		}
	}
	return
}

func main() {
	file, err := os.Open("2021/day22/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	cubes1 := make(map[[3]int]bool)
	cubes := make([]cube, 0)
	xPoints := make([]int, 0)
	yPoints := make([]int, 0)
	zPoints := make([]int, 0)

	var minMax [6]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		onOff := strings.Split(line[0], " ")[0]
		line[0] = strings.Trim(strings.Trim(line[0], "on "), "off ")
		xArea := strings.Split(strings.Trim(line[0], "x="), "..")
		if xmin := utils.ToInt(xArea[0]); xmin < minMax[0] {
			minMax[0] = xmin
		}
		if xmax := utils.ToInt(xArea[1]); xmax > minMax[1] {
			minMax[1] = xmax
		}

		yArea := strings.Split(strings.Trim(line[1], "y="), "..")
		if ymin := utils.ToInt(yArea[0]); ymin < minMax[2] {
			minMax[2] = ymin
		}
		if ymax := utils.ToInt(yArea[0]); ymax > minMax[3] {
			minMax[3] = ymax
		}
		zArea := strings.Split(strings.Trim(line[2], "z="), "..")
		if zmin := utils.ToInt(zArea[0]); zmin < minMax[4] {
			minMax[4] = zmin
		}
		if zmax := utils.ToInt(zArea[0]); zmax > minMax[5] {
			minMax[5] = zmax
		}
		fmt.Println(onOff, xArea, yArea, zArea)
		for x := utils.ToInt(xArea[0]); x <= utils.ToInt(xArea[1]) && utils.IntAbs(x) <= 50; x++ {
			for y := utils.ToInt(yArea[0]); y <= utils.ToInt(yArea[1]) && utils.IntAbs(y) <= 50; y++ {
				for z := utils.ToInt(zArea[0]); z <= utils.ToInt(zArea[1]) && utils.IntAbs(z) <= 50; z++ {

					if onOff == "on" {
						cubes1[[3]int{x, y, z}] = true
					} else {
						cubes1[[3]int{x, y, z}] = false
					}

				}
			}
		}

		xPoints = append(xPoints, utils.ToInt(xArea[0]))
		xPoints = append(xPoints, utils.ToInt(xArea[1])+1)
		yPoints = append(yPoints, utils.ToInt(yArea[0]))
		yPoints = append(yPoints, utils.ToInt(yArea[1])+1)
		zPoints = append(zPoints, utils.ToInt(zArea[0]))
		zPoints = append(zPoints, utils.ToInt(zArea[1])+1)
		newCube := cube{onOff, toIntArea(xArea), toIntArea(yArea), toIntArea(zArea)}
		cubes = append(cubes, newCube)
	}
	sort.Ints(xPoints)
	sort.Ints(yPoints)
	sort.Ints(zPoints)
	xPoints = unique(xPoints)
	yPoints = unique(yPoints)
	zPoints = unique(zPoints)
	counter := 0
	for key, val := range cubes1 {
		if utils.IntAbs(key[0]) <= 50 && utils.IntAbs(key[1]) <= 50 && utils.IntAbs(key[2]) <= 50 {
			if val == true {
				counter++
			}
		}
	}
	fmt.Println(counter)
	fmt.Println(minMax)
	counter = 0
	var xspan, yspan, zspan int
	fmt.Println(len(xPoints), len(yPoints), len(zPoints))
	for xi, x := range xPoints {
		fmt.Println(xi)
		for yi, y := range yPoints {
			for zi, z := range zPoints {
				if findStatus(x, y, z, cubes) {
					if xi != len(xPoints)-1 {
						if findStatus(x+1, y, z, cubes) {
							xspan = utils.IntAbs((xPoints[xi+1]) - x)
						} else {
							xspan = 1
						}
					} else {
						xspan = 1
					}
					if yi != len(yPoints)-1 {
						if findStatus(x, y+1, z, cubes) {
							yspan = utils.IntAbs((yPoints[yi+1]) - y)
						} else {
							yspan = 1
						}
					} else {
						yspan = 1
					}
					if zi != len(zPoints)-1 {
						if findStatus(x, y, z+1, cubes) {
							zspan = utils.IntAbs((zPoints[zi+1]) - z)
						} else {
							zspan = 1
						}
					} else {
						zspan = 1
					}
					counter += xspan * yspan * zspan
				}
			}
		}
	}

	fmt.Println(counter)
}
