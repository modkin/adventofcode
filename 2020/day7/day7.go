package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type multipleBags struct {
	amount int
	color  rune
}

func search_gold(bags []string, possible map[string][]string, found int) int {
	for _, elem := range bags {
		if elem == "shiny gold" {
			found++
		} else {
			found += search_gold(possible[elem], possible, found)
		}
	}
	return found
}

func find_shiny(bags []string, possible map[string][]string) bool {
	tmp := false
	for _, elem := range bags {
		if elem == "shiny gold" {
			tmp = true
		} else {
			tmp = tmp || find_shiny(possible[elem], possible)
		}
	}
	return tmp
}

func main() {
	file, err := os.Open("2020/day7/input")
	if err != nil {
		panic(err)
	}

	possbile := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " bags contain ")
		bag := line[0]
		contents := strings.Split(line[1], ",")
		inside := make([]string, 0)
		for _, c := range contents {
			tmp := strings.Split(strings.TrimSpace(c), " ")
			if tmp[1] != "other" {
				inside = append(inside, tmp[1]+" "+tmp[2])
			}
		}
		possbile[bag] = inside
	}
	//fmt.Println(possbile)
	all_bags := make([]string, 0)
	for keys, _ := range possbile {
		all_bags = append(all_bags, keys)
	}
	//fmt.Println(all_bags)
	total_red := 0
	for _, content := range possbile {
		//if search_gold(possbile[color],possbile,0) > 0 {
		//	total_red++
		//	fmt.Println(color)
		//}
		if find_shiny(content, possbile) {
			total_red++
		}
	}
	fmt.Println(total_red)
}
