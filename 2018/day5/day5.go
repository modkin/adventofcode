package day5

import (
	"io/ioutil"
	"math"
)

func polymerLen(ignore byte) int {

	content, err := ioutil.ReadFile("day5/day5input.txt")
	if err != nil {
		panic(err)
	}
	result := []byte{content[len(content)-1]}
	content = content[:len(content)-1]

	for _, elem := range content {
		if ignore == elem || ignore+32 == elem {
			continue
		} else if math.Abs(float64(result[len(result)-1])-float64(elem)) == 32 {
			result = result[:len(result)-1]
		} else {
			result = append(result, elem)
		}
	}
	return len(result) - 1
}

func Task1() int {
	return polymerLen(0)
}

func Task2() int {
	shortestPolymer := math.MaxInt16
	for i := 65; i < 91; i++ {
		polylen := polymerLen(byte(i))
		if polylen < shortestPolymer {
			shortestPolymer = polylen
		}
	}
	return shortestPolymer
}
