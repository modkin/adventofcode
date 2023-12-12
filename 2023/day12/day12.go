package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func damageReport(input string) (report []int) {
	reg := regexp.MustCompile(`#+`)
	found := reg.FindAllString(input, -1)
	for _, s := range found {
		report = append(report, len(s))
	}
	return
}

func compareInt(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	same := true
	for i, i2 := range left {
		if i2 != right[i] {
			same = false
		}
	}
	return same
}

func flip(input string) (string, bool) {
	if input == "#" {
		return ".", true
	} else {
		return "#", false
	}
}

var cache map[string]int

func countArrangements(leftString string, leftDamage []int) int {
	hashed := leftString
	for _, i2 := range leftDamage {
		hashed += string(rune(i2))
	}
	if val, ok := cache[hashed]; ok {
		return val
	}
	if len(leftDamage) == 0 {
		if strings.Count(leftString, "#") > 0 {
			return 0
		} else {
			return 1
		}
	}
	if strings.Count(leftString, ".") == len(leftString) {
		return 0
	}

	nextSize := leftDamage[0]
	firstqm := math.MaxInt
	if ret := regexp.MustCompile(`\?`).FindStringIndex(leftString); len(ret) != 0 {
		firstqm = ret[0]
	}
	firsthash := math.MaxInt
	if ret := regexp.MustCompile(`#`).FindStringIndex(leftString); len(ret) != 0 {
		firsthash = ret[0]
	}
	if firsthash < firstqm {
		end := len(leftString)
		if firsthash+nextSize < end {
			end = firsthash + nextSize
		}
		startingHashCount := strings.Count(strings.Replace(leftString, "?", "#", -1)[firsthash:end], "#")
		if startingHashCount != nextSize {
			return 0
		}

		if strings.Count(strings.Replace(leftString, "?", "#", -1), "#") == nextSize && len(leftDamage) == 1 {
			return 1
		}
		if len(leftString) <= firsthash+nextSize {
			return 0
		}
		if leftString[firsthash+nextSize] == '.' || leftString[firsthash+nextSize] == '?' {
			ret := countArrangements(leftString[firsthash+nextSize+1:], leftDamage[1:])
			cache[hashed] = ret
			return ret
		}
	} else {
		changedString1 := strings.Replace(leftString, "?", "#", 1)
		changedString2 := strings.Replace(leftString, "?", ".", 1)
		ret := countArrangements(changedString1, leftDamage) + countArrangements(changedString2, leftDamage)
		cache[hashed] = ret
		return ret
	}
	return 0
}

func main() {
	cache = make(map[string]int)
	file, err := os.Open("2023/day12/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var springs [][]string
	var damage [][]int
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	multi := 4

	for _, line := range lines {
		split := strings.Fields(line)
		springString := split[0]
		damageString := split[1]
		splitPart0 := split[0]
		splitPart1 := split[1]
		for i := 0; i < multi; i++ {
			springString = strings.Join([]string{springString, splitPart0}, "?")
			damageString = strings.Join([]string{damageString, splitPart1}, ",")
		}

		var newSprings []string
		for _, s := range strings.Split(springString, "") {
			newSprings = append(newSprings, s)
		}
		springs = append(springs, newSprings)
		var newDamage []int
		for _, s2 := range strings.Split(damageString, ",") {
			newDamage = append(newDamage, utils.ToInt(s2))
		}
		damage = append(damage, newDamage)
	}

	_ = func() int {
		sumArrangements := 0
		for i, spring := range springs {
			arrangments := 0
			reg := regexp.MustCompile(`\?`)
			joinedString := strings.Join(spring, "")
			found := reg.FindAllStringIndex(joinedString, -1)
			joinedString = strings.ReplaceAll(joinedString, "?", ".")
			if compareInt(damageReport(joinedString), damage[i]) {
				arrangments++
			}
			totalBroken := utils.SumSlice(damage[i])
			spring = strings.Split(joinedString, "")
			for j := 0; j < int(math.Exp2(float64(len(found))))-1; j++ {
				over := false
				for i2, idxArray := range found {
					idx := idxArray[0]

					if i2 == 0 {
						spring[idx], over = flip(spring[idx])
					} else {
						if over {
							spring[idx], over = flip(spring[idx])
						}
					}

				}
				if strings.Count(strings.Join(spring, ""), "#") == totalBroken {
					if compareInt(damageReport(strings.Join(spring, "")), damage[i]) {
						arrangments++
					}
				}
			}
			sumArrangements += arrangments
			fmt.Println(i, arrangments)
		}
		return sumArrangements
	}

	totalSum := 0
	for i, spring := range springs {
		sum := countArrangements(strings.Join(spring, ""), damage[i])
		fmt.Println(i, sum)
		totalSum += sum
	}
	fmt.Println(totalSum)

}
