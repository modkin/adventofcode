package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strings"
)

type network struct {
	pcs map[string]bool
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day23/input")

	//allNetworks := []network{}
	connections := [][2]string{}

	threes := make(map[[3]string]bool)
	for _, line := range lines {
		s := strings.Split(line, "-")
		one := s[0]
		two := s[1]
		sli := []string{one, two}
		slices.Sort(sli)
		connections = append(connections, [2]string{sli[0], sli[1]})
		//notThere := true
		//for _, net := range allNetworks {
		//	for i2, _ := range net.pcs {
		//		if one == i2 || two == i2 {
		//			notThere = false
		//			net.pcs[one] = true
		//			net.pcs[two] = true
		//		}
		//	}
		//}
		//if notThere {
		//	allNetworks = append(allNetworks, network{map[string]bool{one: true, two: true}})
		//}
	}
	for _, c1 := range connections {
		for _, c2 := range connections {
			if c1 == c2 {
				continue
			}
			if c1[0] == c2[0] || c1[0] == c2[1] {
				var missingCon [2]string
				if c1[0] == c2[0] {
					missingCon = [2]string{c1[1], c2[1]}
				} else {
					missingCon = [2]string{c1[1], c2[0]}
				}

				for _, c3 := range connections {
					if c3 == missingCon {
						m := map[string]bool{c1[0]: true, c1[1]: true, c2[0]: true, c2[1]: true, c3[0]: true, c3[1]: true}
						s := []string{}
						for i3, _ := range m {
							s = append(s, i3)
						}
						slices.Sort(s)
						if len(s) == 3 {
							threes[[3]string{s[0], s[1], s[2]}] = true
						}
					}

				}

				//allNetworks = append(allNetworks, network{map[string]bool{c1[0]: true, c1[1]: true, c2[0]: true, c2[1]: true}})
			}
		}
	}
	sum := 0
	fmt.Println(len(threes))
	for i, _ := range threes {
		for _, p := range i {
			if p[0] == 't' {
				sum += 1
				break
			}
		}
	}
	fmt.Println(sum)
}
