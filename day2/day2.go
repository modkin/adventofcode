package day2

import (
	"bufio"
	"os"
	"strings"
)

func FindSpecial(name string) string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var words [][]string
	for scanner.Scan() {
		word := scanner.Text()
		chars := strings.Split(word, "")
		words = append(words, chars)
	}

	for _, element := range words {
		for _, element2 := range words {
			var matches = 0
			for index, char1 := range element {
				if char1 == element2[index] {
					matches++
				}
			}
			if matches == len(element)-1 {
				for index, char := range element {
					if char != element2[index] {
						result := append(element[0:index], element[index+1:]...)
						return strings.Join(result, "")
					}
				}
			}

		}
	}
	return "Not Found"
}

func Checksum(name string) int {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var threeChars = 0
	var twoChars = 0

	for scanner.Scan() {
		word := scanner.Text()
		chars := strings.Split(word, "")
		m := map[string]bool{}
		twofound := true
		threefound := true
		for _, element := range chars {
			m[element] = true
		}
		for key, _ := range m {
			charcount := strings.Count(word, key)
			if charcount == 2 && twofound {
				twoChars++
				twofound = false
			}
			if charcount == 3 && threefound {
				threeChars++
				threefound = false
			}
		}
	}
	return threeChars * twoChars
}
