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

func main() {
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
	for _, line := range lines {
		split := strings.Fields(line)
		var newSprings []string
		for _, s := range strings.Split(split[0], "") {
			newSprings = append(newSprings, s)
		}
		springs = append(springs, newSprings)
		var newDamage []int
		for _, s2 := range strings.Split(split[1], ",") {
			newDamage = append(newDamage, utils.ToInt(s2))
		}
		damage = append(damage, newDamage)
	}

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
			if compareInt(damageReport(strings.Join(spring, "")), damage[i]) {
				arrangments++
			}
		}
		sumArrangements += arrangments
		//fmt.Println(i, arrangments)
	}

	fmt.Println(sumArrangements)
}
