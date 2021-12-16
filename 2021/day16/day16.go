package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func toInt64(input []string) int64 {
	number, err := strconv.ParseInt(strings.Join(input, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	return number
}

func onlyZero(input []string) bool {
	isZero := true
	for _, i := range input {
		if i != "0" {
			isZero = false
		}
	}
	return isZero
}

func parse(input []string, stopPos int64, stopCount int64) (versions []int64, literals []int64, pos int64) {
	count := int64(0)
	for pos <= stopPos && count <= stopCount {
		if onlyZero(input[pos:]) {
			break
		}
		version := toInt64(input[pos : pos+3])
		versions = append(versions, version)
		typeID := toInt64(input[pos+3 : pos+6])
		pos += 6
		switch typeID {
		case 4:
			number := make([]string, 0)
			for {
				group := input[pos : pos+5]
				number = append(number, group[1:]...)
				pos += 5
				if group[0] == "0" {
					break
				}
			}
			numberInt, _ := strconv.ParseInt(strings.Join(number, ""), 2, 64)
			literals = append(literals, numberInt)
		default:
			lengthTypeID := input[pos]
			var length int64
			if lengthTypeID == "0" {
				length = toInt64(input[pos+1 : pos+15])
				pos += 1 + 15
				stop := pos + length
				newVersions, newLiterals, endPos := parse(input[pos:], stop, math.MaxInt64)
				versions = append(versions, newVersions...)
				literals = append(literals, newLiterals...)
				pos += endPos
			} else {
				length = toInt64(input[pos+1 : pos+11])
				pos += 1 + 11
				newVersions, newLiterals, endPos := parse(input[pos:], math.MaxInt64, length)
				versions = append(versions, newVersions...)
				literals = append(literals, newLiterals...)
				pos += endPos
			}
		}
	}
	return
}

func main() {
	file, err := os.Open("2021/day16/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	asci := make(map[string][]string)
	asci["0"] = strings.Split("0000", "")
	asci["1"] = strings.Split("0001", "")
	asci["2"] = strings.Split("0010", "")
	asci["3"] = strings.Split("0011", "")
	asci["4"] = strings.Split("0100", "")
	asci["5"] = strings.Split("0101", "")
	asci["6"] = strings.Split("0110", "")
	asci["7"] = strings.Split("0111", "")
	asci["8"] = strings.Split("1000", "")
	asci["9"] = strings.Split("1001", "")
	asci["A"] = strings.Split("1010", "")
	asci["B"] = strings.Split("1011", "")
	asci["C"] = strings.Split("1100", "")
	asci["D"] = strings.Split("1101", "")
	asci["E"] = strings.Split("1110", "")
	asci["F"] = strings.Split("1111", "")

	code := make([]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for _, l := range line {
			code = append(code, asci[l]...)
		}
	}
	//fmt.Println(code)
	versions, literals, _ := parse(code, int64(len(code)), math.MaxInt64)
	fmt.Println(versions, literals)
	sum := int64(0)
	for _, version := range versions {
		sum += version
	}
	fmt.Println("Day 16.1:", sum)
}
