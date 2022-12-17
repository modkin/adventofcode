package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checksum(bottom string, shapeID, windID int) string {
	return bottom + "-" + strconv.Itoa(shapeID) + "-" + strconv.Itoa(windID)
}

func countRow(game map[[2]int]string, ymax int) int {
	counter := 0
	for i := 0; i < 7; i++ {
		if _, ok := game[[2]int{i, ymax}]; ok {
			counter++
		}
	}
	return counter
}
func main() {
	shapes := make([][][2]int, 0)
	file, err := os.Open("2022/day17/shapes")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	shape := make([][2]int, 0)
	y := 1
	shapeHeight := []int{1, 3, 3, 4, 2}
	count := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			shapes = append(shapes, shape)
			shape = make([][2]int, 0)
			y = 1
			count++
		} else {
			split := strings.Split(scanner.Text(), "")
			for i, s := range split {
				if s == "#" {
					shape = append(shape, [2]int{i, shapeHeight[count] - y})
				}
			}
			y++
		}
	}
	shapes = append(shapes, shape)

	//for _, i2 := range shapes {
	//	utils.Print2D2IntGrid(i2)
	//}

	file, err = os.Open("2022/day17/input")

	if err != nil {
		panic(err)
	}

	wind := make([]string, 0)
	game := make(map[[2]int]string)
	seenBefore := make(map[string][2]int)
	target := 1000000000000

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		wind = strings.Split(scanner.Text(), "")

	}
	curShape := 0
	xOff := 2
	yOff := 3
	yMax := 0
	windDir := 0
	windPos := 0
	for i := 0; i < 7; i++ {
		game[[2]int{i, -1}] = "@"
	}
	rockCounter := 0
	//seenBefore[checksum(countRow(game, yMax), curShape, windPos)] = rockCounter
	for {
		w := wind[windPos]
		if w == ">" {
			windDir = 1
		} else {
			windDir = -1
		}
		windPush := true
		for _, i := range shapes[curShape] {
			if i[0]+xOff+windDir >= 7 || i[0]+xOff+windDir < 0 {
				windPush = false
			} else {
				if _, ok := game[[2]int{i[0] + xOff + windDir, i[1] + yOff}]; ok {
					windPush = false
				}
			}
		}
		if windPush {
			xOff += windDir
		}
		windPos = (windPos + 1) % len(wind)

		//utils.PrintTetris(game, shapes[curShape], xOff, yOff)

		stopps := false
		for _, i := range shapes[curShape] {
			if _, ok := game[[2]int{i[0] + xOff, i[1] + yOff - 1}]; ok {
				stopps = true
			}
		}
		if stopps {
			for _, i := range shapes[curShape] {
				game[[2]int{i[0] + xOff, i[1] + yOff}] = "#"
				if i[1]+yOff > yMax {
					yMax = i[1] + yOff
				}
			}
			yOff = yMax + 4
			xOff = 2
			curShape = (curShape + 1) % len(shapes)
			rockCounter++
			if rockCounter == target {
				fmt.Println("Day 17.1", yMax+1)
				break
			}
			count = countRow(game, yMax)
			if count >= 6 {
				bottomRow := ""
				for i := 0; i < 7; i++ {
					if _, ok := game[[2]int{i, yMax}]; ok {
						bottomRow += "#"
					} else {
						bottomRow += "."
					}
				}
				if old, ok := seenBefore[checksum(bottomRow, curShape, windPos)]; ok {
					fmt.Println("WUHU:", old, rockCounter, rockCounter-old[0], yMax-old[1])
					seenBefore[checksum(bottomRow, curShape, windPos)] = [2]int{rockCounter, yMax}
					rockDiff := rockCounter - old[0]
					yDiff := yMax - old[1]
					yMaxOld := yMax
					for rockCounter+rockDiff < target {
						rockCounter += rockDiff
						yMax += yDiff
					}

					for i := 0; i < 7; i++ {
						if val, ok2 := game[[2]int{i, yMaxOld}]; ok2 {
							game[[2]int{i, yMax}] = val
						} else {
							yfoo := 0
							for {
								if _, ok3 := game[[2]int{i, yMaxOld - yfoo}]; ok3 {
									fmt.Println("MUUU", yfoo)
									break
								} else {
									yfoo++
								}
							}
						}
					}
					yOff = yMax + 4
					fmt.Println(rockCounter, yMax, yOff)
					//utils.PrintTetris(game, shapes[curShape], xOff, yOff)
					//break
				} else {
					seenBefore[checksum(bottomRow, curShape, windPos)] = [2]int{rockCounter, yMax}
				}
			}
			//utils.PrintTetris(game, shapes[curShape], xOff, yOff)
		} else {
			yOff--
		}
	}
}
