package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(left [2]int, right [2]int) [2]int {
	return [2]int{left[0] + right[0], left[1] + right[1]}
}

type position struct {
	pos  [2]int
	cost int
}

func countInOffset(offset [2]int, allPos map[[2][2]int]int) int {
	counter := 0
	for i, _ := range allPos {
		if i[1] == offset {
			counter++
		}
	}
	return counter
}

var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func findPossiblePos(start [2]int, garden map[[2]int]string, numSteps int, maxX int, maxY int) int {
	allPos := map[[2]int]bool{start: true}
	for i := 0; i < numSteps; i++ {
		newAllPos := make(map[[2]int]bool)
		for pos, _ := range allPos {
			for _, dir := range dirs {
				newPos := sum(pos, dir)
				if newPos[0] < 0 {
					continue
				}
				if newPos[0] > maxX {
					continue
				}
				if newPos[1] < 0 {
					continue
				}
				if newPos[1] > maxY {
					continue
				}
				if garden[newPos] != "#" {
					newAllPos[newPos] = true
				}
			}
		}
		allPos = newAllPos
	}
	return len(allPos)
}

func main() {
	file, err := os.Open("2023/day21/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	garden := make(map[[2]int]string)
	y := 0
	var maxX int
	var maxY int
	var start [2]int
	for scanner.Scan() {
		for x, cost := range scanner.Text() {
			garden[[2]int{x, y}] = string(cost)
			if string(cost) == "S" {
				start = [2]int{x, y}
			}
		}
		y++
		maxX = len(scanner.Text()) - 1

	}
	maxY = y - 1
	fmt.Println("maxX:", maxX, "maxY:", maxY)
	sideLength := maxX + 1

	//utils.Print2DStringsGrid(garden)
	fmt.Println("Day 21.1 1:", findPossiblePos(start, garden, 64, maxX, maxY))
	fmt.Println(maxX/2 + 1)

	lowerCount := 0
	higherCount := 0
	for i := 129; true; i++ {
		tmpCount := findPossiblePos(start, garden, i, maxX, maxY)
		if tmpCount == lowerCount {
			fmt.Println(i)
			break

		}
		lowerCount = higherCount
		higherCount = tmpCount
	}
	fmt.Println(lowerCount, higherCount)
	fmt.Println("uneven:", findPossiblePos(start, garden, 305, maxX, maxY))
	//fmt.Println("Day 21.1 1:", findPossiblePos(start, garden, 1000, maxX, maxY))
	//fmt.Println("Day 21.1 1:", findPossiblePos(start, garden, 1001, maxX, maxY))
	//fmt.Println("Day 21.1 1:", findPossiblePos(start, garden, 1002, maxX, maxY))

	//start = [2]int{0, 10}
	allPos := make(map[[2][2]int]int)
	allPos[[2][2]int{start, [2]int{0, 0}}] = 0

	//var allPos []position
	//allPos = append(allPos, position{
	//	start, 0,
	//})
	//var dirs [][2]int
	innerOffsets := make(map[[2]int]int)
	enterNewOffsets := make(map[[2]int]int)
	//visited := make(map[[2][2]int]bool)
	//dirs = append(dirs, [2]int{1, 0})
	//dirs = append(dirs, [2]int{-1, 0})
	//dirs = append(dirs, [2]int{0, 1})
	//dirs = append(dirs, [2]int{0, -1})
	//reachable := make(map[[2]int]bool)
	numSteps := 26501365
	numSteps = 300
	maxFillSteps := 221
	//maxFillSteps = 41
	for i := 1; i <= numSteps; i++ {
		fmt.Println(i)
		newAllPos := make(map[[2][2]int]int)
		//var newAllPos []position
		for pos, oldDist := range allPos {
			for _, dir := range dirs {

				oldOffset := pos[1]
				newoffset := oldOffset
				newPos := sum(pos[0], dir)
				//if utils.IntAbs(newoffset[0]) != 0 && utils.IntAbs(newoffset[1]) != 0 {
				//	continue
				//}
				if newPos[0] < 0 {
					newPos[0] = maxX
					newoffset[0] -= 1
				}
				if newPos[0] > maxX {
					newPos[0] = 0
					newoffset[0] += 1
				}
				if newPos[1] < 0 {
					newPos[1] = maxY
					newoffset[1] -= 1
				}
				if newPos[1] > maxY {
					newPos[1] = 0
					newoffset[1] += 1
				}
				newPosition := [2][2]int{newPos, newoffset}

				_, ok := innerOffsets[newoffset]
				if garden[newPos] != "#" && !ok {
					if countInOffset(newoffset, allPos) == higherCount {
						innerOffsets[newoffset] = i
						if newoffset[0] == 0 && newoffset[1] < 1 {
							//fmt.Println(newoffset, i-enterNewOffsets[newoffset])
						}
						if tmp := i - enterNewOffsets[newoffset]; tmp > maxFillSteps {
							maxFillSteps = tmp
						}
						break
					} else {
						newAllPos[newPosition] = oldDist + 1
					}
					if _, ok2 := enterNewOffsets[newoffset]; !ok2 {
						enterNewOffsets[newoffset] = i
						//if newoffset[0] == 0 && newoffset[1] > 0 {
						//fmt.Println("New offset:", newoffset, newPos, i-enterNewOffsets[oldOffset], i)
						//}
					}

				}

			}
		}
		allPos = newAllPos
		//for ints, _ := range innerOffsets {
		//	fmt.Println("offset: ", ints)
		//	fmt.Println("count: ", countInOffset(ints))
		//}
	}

	//quadStartIndex := 132

	quadrantArea := func(offsetStart [2]int, quadStart [2]int) int {
		topRightStartIndex := maxX + 2 //enterNewOffsets[offsetStart]
		fullMaps := (numSteps - topRightStartIndex - maxFillSteps) / sideLength
		firstNotFullIndex := sideLength*fullMaps + topRightStartIndex
		//stepsleft := (numSteps - topRightStartIndex) % maxX
		//fmt.Println(topRightStartIndex, fullMaps, firstNotFullIndex)

		firstCount := findPossiblePos(quadStart, garden, 301, maxX, maxY)
		secondCount := findPossiblePos(quadStart, garden, 302, maxX, maxY)
		if (numSteps-topRightStartIndex)%2 == 2 {
			firstCount, secondCount = secondCount, firstCount
		}

		areaSum := 0
		for i := 1; i <= fullMaps; i += 2 {
			areaSum += i * firstCount
			if (i + 1) <= fullMaps {
				areaSum += (i + 1) * secondCount
			}
		}
		amount := fullMaps + 1
		for i := firstNotFullIndex; i <= numSteps; i += sideLength {
			areaSum += findPossiblePos(quadStart, garden, numSteps-i, maxX, maxY) * amount
			amount++
		}
		return areaSum
	}

	rectArea := func(recStart [2]int) int {
		topRightStartIndex := start[0] + 1
		fullMaps := (numSteps - topRightStartIndex - maxFillSteps) / sideLength
		firstNotFullIndex := sideLength*fullMaps + topRightStartIndex
		//stepsleft := (numSteps - topRightStartIndex) % maxX
		//fmt.Println(topRightStartIndex, fullMaps, firstNotFullIndex)

		areaSum := 0
		firstCount := findPossiblePos(recStart, garden, 301, maxX, maxY)
		secondCount := findPossiblePos(recStart, garden, 302, maxX, maxY)
		if (numSteps-topRightStartIndex)%2 == 2 {
			firstCount, secondCount = secondCount, firstCount
		}
		for i := 1; i <= fullMaps; i += 2 {
			areaSum += firstCount
			if (i + 1) <= fullMaps {
				areaSum += secondCount
			}
		}
		for i := firstNotFullIndex; i <= numSteps; i += sideLength {
			areaSum += findPossiblePos(recStart, garden, numSteps-i, maxX, maxY)
		}
		return areaSum
	}
	//fmt.Println("Part 1:", findPossiblePos(start, garden, 6))

	//for i := 0; i < 23; i++ {
	//	foo := findPossiblePos([2]int{0, maxY}, garden, i, maxX, maxY)
	//	fmt.Println("i :", i, "area:", foo)
	//}
	innerCounter := 0
	innerCounterTopLeft := 0
	for off, i := range innerOffsets {
		if off[0] != 0 && off[1] != 0 {
			if (numSteps-i)%2 == 1 {
				innerCounterTopLeft += higherCount
			} else {
				innerCounterTopLeft += lowerCount
			}
		} else {
			if (numSteps-i)%2 == 1 {
				innerCounter += higherCount
			} else {
				innerCounter += lowerCount
			}
		}
	}
	posCounter := 0
	posCounterTopLeft := 0
	tmpCounter := 0
	tmpCounter2 := 0
	for i, _ := range allPos {
		if i[1][0] != 0 && i[1][1] != 0 {
			posCounterTopLeft++
		} else {
			posCounter++
		}
		if i[1] == [2]int{1, 0} {
			tmpCounter++
		}
		if i[1] == [2]int{2, 0} {
			tmpCounter2++
		}
	}

	fmt.Println("GO")
	areaSum := quadrantArea([2]int{1, -1}, [2]int{0, maxY})
	fmt.Println(areaSum)
	areaSum += quadrantArea([2]int{1, 1}, [2]int{0, 0})
	fmt.Println(areaSum)
	areaSum += quadrantArea([2]int{-1, -1}, [2]int{maxX, maxY})
	fmt.Println(areaSum)
	areaSum += quadrantArea([2]int{-1, 1}, [2]int{maxX, 0})
	fmt.Println(areaSum)
	areaSum += rectArea([2]int{0, (maxX / 2) + 1})
	fmt.Println(areaSum)
	areaSum += rectArea([2]int{(maxX / 2) + 1, 0})
	fmt.Println(areaSum)
	areaSum += rectArea([2]int{maxX, (maxX / 2) + 1})
	fmt.Println(areaSum)
	areaSum += rectArea([2]int{(maxX / 2) + 1, maxY})
	fmt.Println(areaSum)
	areaSum += findPossiblePos(start, garden, numSteps%100+200, maxX, maxY)
	fmt.Println("tmp:", tmpCounter, "tmp2:", tmpCounter2)
	fmt.Println("New:", posCounter+innerCounter+areaSum)
	fmt.Println("Old:", posCounterTopLeft+innerCounterTopLeft+posCounter+innerCounter)
	fmt.Println(posCounterTopLeft, innerCounterTopLeft, posCounter, innerCounter)
	fmt.Println(posCounter, innerCounter, areaSum)
	//fmt.Println(len(visited))

}
