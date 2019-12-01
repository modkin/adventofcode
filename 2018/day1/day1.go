package day1

import (
	"bufio"
	"os"
	"strconv"
)

func FindFirstTwice(name string) int {

	var m = make(map[int]bool)
	var sum = 0

	for true {
		file, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			if m[sum] {
				return sum
			}
			m[sum] = true
			sum += i

		}
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}
	return 0
}

func SumFile(name string) int {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var sum = 0
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		sum += i
	}
	return sum
}
