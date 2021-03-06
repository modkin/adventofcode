package day3

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type claim struct {
	claimId int
	xOffset int
	yOffset int
	xSize   int
	ySize   int
}

func toInt(str string) int {
	ret, _ := strconv.Atoi(str)
	return ret
}

func findNonOverlapClaim(fabric *[1000][1000]int, claim claim) (bool, int) {
	for i := claim.xOffset; i < claim.xOffset+claim.xSize; i++ {
		for j := claim.yOffset; j < claim.yOffset+claim.ySize; j++ {
			if fabric[i][j] != 1 {
				return false, 0
			}
		}
	}
	return true, claim.claimId
}

func fillFabric(fabric *[1000][1000]int, claim claim) {
	for i := claim.xOffset; i < claim.xOffset+claim.xSize; i++ {
		for j := claim.yOffset; j < claim.yOffset+claim.ySize; j++ {
			fabric[i][j]++
		}
	}
}

func countConflicts(fabric *[1000][1000]int) (count int) {
	count = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				count++
			}
		}
	}
	return
}

func createFabricWithConflicts() ([1000][1000]int, []claim) {
	file, err := os.Open("day3/day3input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	fabric := [1000][1000]int{}
	var claimList []claim

	re := regexp.MustCompile("#([0-9]*) @ ([0-9]*),([0-9]*): ([0-9]*)x([0-9]*)")
	for scanner.Scan() {
		word := scanner.Text()
		result := re.FindAllStringSubmatch(word, -1)
		v := claim{toInt(result[0][1]),
			toInt(result[0][2]), toInt(result[0][3]),
			toInt(result[0][4]), toInt(result[0][5])}
		claimList = append(claimList, v)
	}

	for _, element := range claimList {
		fillFabric(&fabric, element)
	}

	return fabric, claimList
}

func Task1() int {
	fabric, _ := createFabricWithConflicts()
	return countConflicts(&fabric)
}

func Task2() int {
	fabric, claimList := createFabricWithConflicts()
	found := false
	id := 0

	for _, element := range claimList {
		found, id = findNonOverlapClaim(&fabric, element)
		if found {
			break
		}
	}
	return id
}
