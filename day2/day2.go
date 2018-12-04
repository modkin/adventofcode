package day2

import (
	"bufio"
	"os"
	"strings"
)

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
