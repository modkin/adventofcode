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

func sorted2(one, two string) [2]string {
	str := []string{one, two}
	slices.Sort(str)
	return [2]string{str[0], str[1]}
}
func sorted3(one, two, three string) [3]string {
	str := []string{one, two, three}
	slices.Sort(str)
	return [3]string{str[0], str[1], str[2]}
}

func removeDup(in [][]string) [][]string {
	tmpMap := make(map[string]bool)
	for _, s := range in {
		str := strings.Join(s, ",")
		tmpMap[str] = true
	}
	ret := [][]string{}
	for s, _ := range tmpMap {
		ret = append(ret, strings.Split(s, ","))
	}
	return ret
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day23/input")

	//allNetworks := []network{}
	connections := [][2]string{}
	connectionMap := make(map[[2]string]bool)
	pcMap := make(map[string]bool)

	//threes := make(map[[3]string]bool)
	for _, line := range lines {
		s := strings.Split(line, "-")
		one := s[0]
		two := s[1]
		sli := []string{one, two}
		pcMap[one] = true
		pcMap[two] = true
		slices.Sort(sli)
		connections = append(connections, [2]string{sli[0], sli[1]})
		connectionMap[[2]string{sli[0], sli[1]}] = true

	}

	threeNets := make(map[[3]string]bool)
	for _, con := range connections {
		for pc, _ := range pcMap {
			if con[0] != pc && con[1] != pc {
				if connectionMap[sorted2(con[0], pc)] && connectionMap[sorted2(con[1], pc)] {
					threeNets[sorted3(con[0], con[1], pc)] = true
				}
			}
		}
	}
	part1 := 0
	for i, _ := range threeNets {
		for _, p := range i {
			if p[0] == 't' {
				part1 += 1
				break
			}
		}
	}
	fmt.Println("Day 23.1:", part1)

	longestNets := [][]string{}
	for net, _ := range threeNets {
		longestNets = append(longestNets, []string{net[0], net[1], net[2]})
	}
	foundLonger := true
	for foundLonger {
		foundLonger = false
		newLongestNets := [][]string{}
	outer:
		for _, net := range longestNets {
			for pc, _ := range pcMap {
				allCon := true
				for _, pc2 := range net {
					if !connectionMap[sorted2(pc, pc2)] {
						allCon = false
					}
				}
				if allCon {
					newNet := append(net, pc)
					slices.Sort(newNet)
					newLongestNets = append(newLongestNets, newNet)
					continue outer
				}
			}
		}
		newLongestNets = removeDup(newLongestNets)
		if len(newLongestNets) > 0 {
			longestNets = newLongestNets
			foundLonger = true
		}
	}
	var str []string
	for _, s := range longestNets[0] {
		str = append(str, s)
	}
	slices.Sort(str)
	fmt.Println("Day 23.2:", strings.Join(str, ","))
}
