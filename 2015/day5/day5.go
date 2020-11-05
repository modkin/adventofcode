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
	file2, err := os.Open("2015/day5/input")
	if err != nil {
		panic(err)
	}
	scanner = bufio.NewScanner(file2)
	nice_strings = 0
	for scanner.Scan() {
		word := scanner.Text()
		double_strings := make(map[[2]byte]int)
		wordByte := []byte(word)
		foundDouble := false
		foundRepeatWithBetween := false
		for i := 0; i < len(wordByte); i++ {
			if i < len(wordByte)-1 {
				newString := [2]byte{wordByte[i], wordByte[i+1]}
				if idx, ok := double_strings[newString]; ok {
					if idx < i-1 {
						foundDouble = true
					}
				} else {
					double_strings[newString] = i
				}
			}
			if i > 1 {
				if wordByte[i] == wordByte[i-2] {
					foundRepeatWithBetween = true
				}
			}

		}

		if foundDouble && foundRepeatWithBetween {
			nice_strings += 1
		}
	}
	fmt.Println("Task 5.2:", nice_strings)
}
