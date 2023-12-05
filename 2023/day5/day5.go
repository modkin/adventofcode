package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("2023/day5/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var seedList []int
	var seedList2 []int
	newSeedList := make([]int, 0)
	allreadyMapped := make([]int, 0)
	mappingMap := make([][][]int, 0)
	currentMap := -1
	var mapStrings []string
	for scanner.Scan() {
		if len(seedList) == 0 {
			split := strings.Fields(scanner.Text())
			for _, s := range split[1:] {
				seedList = append(seedList, utils.ToInt(s))
				seedList2 = append(seedList2, utils.ToInt(s))
			}
			//for i := 1; i < len(split); i += 2 {
			//	for j := 0; j < utils.ToInt(split[i+1]); j++ {
			//		seedList = append(seedList, utils.ToInt(split[i])+j)
			//	}
			//}
		} else if strings.Contains(scanner.Text(), "map") || scanner.Text() == "" {
			if len(newSeedList) != 0 {
				for _, i2 := range seedList {
					if !utils.IntSliceContains(allreadyMapped, i2) {
						newSeedList = append(newSeedList, i2)
					}
				}
				seedList = utils.CopyIntSlice(newSeedList)
				allreadyMapped = make([]int, 0)
				newSeedList = make([]int, 0)
			}
			if scanner.Text() != "" {
				currentMap++
				mappingMap = append(mappingMap, make([][]int, 0))
				mapStrings = append(mapStrings, scanner.Text())
			}
		} else {
			split := strings.Fields(scanner.Text())
			dst := utils.ToInt(split[0])
			source := utils.ToInt(split[1])
			span := utils.ToInt(split[2])

			//for i := dst; i < dst+span; i++ {
			for _, item := range seedList {
				if item > source && item < source+span && !utils.IntSliceContains(allreadyMapped, item) {
					newSeedList = append(newSeedList, dst+(item-source))
					allreadyMapped = append(allreadyMapped, item)
				} else {
					//newSeedList = append(newSeedList, item)
				}
			}
			//}
			//seedList = utils.CopyIntSlice(newSeedList)
			//if _, ok := mappingMap[currentMap]; !ok {
			//	mappingMap[currentMap] = make([][]int, 0)
			//}
			mappingMap[currentMap] = append(mappingMap[currentMap], []int{dst, source, span})
		}

	}
	for _, i2 := range seedList {
		if !utils.IntSliceContains(allreadyMapped, i2) {
			newSeedList = append(newSeedList, i2)
		}
	}
	seedList = utils.CopyIntSlice(newSeedList)
	allreadyMapped = make([]int, 0)
	newSeedList = make([]int, 0)

	fmt.Println("Day 5.1:", slices.Min(seedList))
	//fmt.Println(mappingMap)
	counter := 0
	for mapIndex, value := range mappingMap {
		fmt.Println(seedList2)
		fmt.Println(mapStrings[mapIndex])
		newSeedList = make([]int, 0)
		leftOver := utils.CopyIntSlice(seedList2)
		for _, ints := range value {
			var newLeftOver []int
			dst := ints[0]
			source := ints[1]
			diff := dst - source
			span := ints[2]

			for i := 0; i != len(leftOver); i += 2 {
				start := leftOver[i]
				itemRange := leftOver[i+1]
				if start > source+span || start+itemRange < source {
					newLeftOver = append(newLeftOver, []int{start, itemRange}...)
				} else {
					if start < source && start+itemRange > source {
						newLeftOver = append(newLeftOver, []int{start, source - start}...)
						itemRange = itemRange - (source - start)
						start = source

					}
					newStart := start + diff
					if start+itemRange > source+span {
						newRange := (source + span) - start
						newSeedList = append(newSeedList, newStart, newRange)
						if itemRange-newRange != 0 {
							newLeftOver = append(newLeftOver, []int{source + span, itemRange - newRange}...)
						}
					} else {
						newSeedList = append(newSeedList, newStart, itemRange)
					}
				}
			}
			leftOver = utils.CopyIntSlice(newLeftOver)
			sum := 0

			for i := 1; i < len(seedList2); i += 2 {
				sum += seedList2[i]
			}
			fmt.Println(sum, counter, slices.Min(seedList2))
			counter++
		}

		seedList2 = utils.CopyIntSlice(newSeedList)
		seedList2 = append(seedList2, leftOver...)
		fmt.Println(seedList2)
	}
	mininmal := math.MaxInt
	for i := 0; i < len(seedList2); i += 2 {
		if seedList2[i] < mininmal {
			mininmal = seedList2[i]
		}
	}
	fmt.Println(mininmal)
}
