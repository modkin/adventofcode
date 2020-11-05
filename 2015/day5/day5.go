package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2015/day5/input")
	if err != nil {
		panic(err)
	}

	vowels := []string{"a", "e", "i", "o", "u"}
	bad_words := []string{"ab", "cd", "pq", "xy"}
	scanner := bufio.NewScanner(file)
	nice_strings := 0
	for scanner.Scan() {
		word := scanner.Text()
		sum_vowels := 0
		for _, vow := range vowels {
			sum_vowels += strings.Count(word, vow)
		}
		if sum_vowels < 3 {
			continue
		}
		word_byte := []byte(word)
		twice := false
		for i := 1; i < len(word_byte); i++ {
			if word_byte[i] == word_byte[i-1] {
				twice = true
			}
		}
		if !twice {
			continue
		}
		is_bad := false
		for _, bad := range bad_words {
			if strings.Contains(word, bad) {
				is_bad = true
			}
		}
		if is_bad {
			continue
		}
		nice_strings += 1
	}
	fmt.Println("Task 5.1:", nice_strings)
}
