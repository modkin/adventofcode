package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func findUncoruppted(input string) int {
	reg := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	found := reg.FindAllStringSubmatch(input, -1)
	result := 0
	for _, s := range found {
		fmt.Println(s[0], s[1], s[2])
		result += utils.ToInt(s[1]) * utils.ToInt(s[2])
	}
	return result
}

func findUncorupptedWithDont(input string, enabled bool) (int, bool) {
	reg := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	reg2 := regexp.MustCompile(`do\(\)|don't\(\)`)
	found := reg.FindAllStringSubmatch(input, -1)
	idxMul := reg.FindAllStringIndex(input, -1)
	idxDo := reg2.FindAllStringIndex(input, -1)
	foundDo := reg2.FindAllString(input, -1)
	result := 0
	//enabled := true
	for i, s := range found {
		multIdx := idxMul[i][0]
		doIdx := -1
		for _, ints := range idxDo {
			if ints[0] < multIdx {
				doIdx++
			}
		}

		if doIdx == -1 {

		} else if foundDo[doIdx] == "do()" {
			enabled = true

			fmt.Println(input[idxDo[doIdx][0]:idxMul[i][1]], "en")
		} else if foundDo[doIdx] == "don't()" {
			enabled = false

			fmt.Println(input[idxDo[doIdx][0]:idxMul[i][1]], "dis")
		} else {
			fmt.Println("PANIC")
		}

		if enabled {
			result += utils.ToInt(s[1]) * utils.ToInt(s[2])
		}

	}
	return result, enabled
}

func main() {
	file, err := os.Open("2024/day3/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result1 := 0
	result2 := 0
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()
		tmpRes, tmpEn := findUncorupptedWithDont(line, enabled)
		result2 += tmpRes
		enabled = tmpEn
		result1 += findUncoruppted(line)
	}
	fmt.Println(result1)
	fmt.Println(result2)
}
