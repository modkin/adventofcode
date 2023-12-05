package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
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
	newSeedList := make([]int, 0)
	allreadyMapped := make([]int, 0)
	for scanner.Scan() {
		if len(seedList) == 0 {
			split := strings.Fields(scanner.Text())
			for _, s := range split[1:] {
				seedList = append(seedList, utils.ToInt(s))
			}
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
}
