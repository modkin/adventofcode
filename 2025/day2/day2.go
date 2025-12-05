package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ranges := strings.Split(utils.ReadFileIntoLines("2025/day2/input")[0], ",")

	invalidSum := 0
	rangesClean := make([]string, 0)
	for _, r := range ranges {
		startStr, endStr, _ := strings.Cut(r, "-")

		if len(startStr) < len(endStr) {
			splitEnd := ""
			for i := 0; i < len(startStr); i++ {
				splitEnd += "9"
			}
			splitStart := strconv.Itoa(utils.ToInt(splitEnd) + 1)

			rangesClean = append(rangesClean, startStr+"-"+splitEnd)
			rangesClean = append(rangesClean, splitStart+"-"+endStr)
		} else {
			rangesClean = append(rangesClean, r)
		}

	}
	invalidSet := make(map[int]bool)
	invalidTwo := 0
	for _, r := range rangesClean {
		startStr, endStr, _ := strings.Cut(r, "-")
		startInt := utils.ToInt(startStr)
		endInt := utils.ToInt(endStr)
		if len(startStr) != len(endStr) {
			//fmt.Println("that was not expected")
			//return
		}

		for numChar := 2; numChar <= len(endStr); numChar++ {

			halfStart := startStr[0 : (len(startStr))/numChar]
			if halfStart == "" {
				halfStart = startStr[0:1]
			}
			halfEnd := endStr[0 : (len(endStr)+1)/numChar]

			//fmt.Println("range: ", startStr, "-", endStr)
			for i := utils.ToInt(halfStart); i <= utils.ToInt(halfEnd); i++ {
				numStr := strconv.Itoa(i)
				for k := 2; k <= numChar; k++ {
					numStr += strconv.Itoa(i)
				}
				num := utils.ToInt(numStr)
				if startInt <= num && num <= endInt {
					if _, ok := invalidSet[num]; !ok {
						invalidSet[num] = true
						invalidSum += num
						if numChar == 2 {
							invalidTwo += num
						}
					}
				}
			}
		}
	}
	fmt.Println("Day 2.1: ", invalidTwo)
	fmt.Println("Day 2.2: ", invalidSum)
}
