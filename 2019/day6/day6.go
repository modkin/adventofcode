package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Planet struct {
	name     string
	orbits   string
	inOrbit  []string
	checksum int
}

func count(planet *Planet, planetMap map[string]Planet, counter int) {
	planet.checksum = counter
	if planet.inOrbit == nil {

	} else {
		for _, other := range planet.inOrbit {
			planet := planetMap[other]
			count(&planet, planetMap, counter+1)
			planetMap[other] = planet
		}
	}
}

func orbitPath(planet Planet, planetMap map[string]Planet, path *[]string) []string {
	if planet.name != "COM" {
		newPath := append(*path, planet.orbits)
		path = &newPath
		orbitPath(planetMap[planet.orbits], planetMap, path)
	} else {
		copy := make([]string, len(*path))
		for i, elem := range *path {
			copy[i] = elem
		}
		return copy
	}
	return nil
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	var planetMap = make(map[string]Planet)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		planets := strings.Split(scanner.Text(), ")")
		if _, exists := planetMap[planets[0]]; exists {
			///WTF go?
			planet := planetMap[planets[0]]
			planet.inOrbit = append(planetMap[planets[0]].inOrbit, planets[1])
			planetMap[planets[0]] = planet
		} else {
			planet := Planet{
				name:     planets[0],
				orbits:   "",
				inOrbit:  []string{planets[1]},
				checksum: 0,
			}
			planetMap[planets[0]] = planet
		}

		if _, exists := planetMap[planets[1]]; exists {
			planet := planetMap[planets[1]]
			if planet.orbits != "" {
				fmt.Println("Error already in orbit")
			}
			planet.orbits = planets[0]
			planetMap[planets[1]] = planet
		} else {
			planet := Planet{
				name:     planets[1],
				orbits:   planets[0],
				inOrbit:  nil,
				checksum: 0,
			}
			planetMap[planets[1]] = planet
		}
	}
	planet := planetMap["COM"]
	count(&planet, planetMap, 0)
	planetMap["COM"] = planet
	sum := 0
	for _, planet := range planetMap {
		sum += planet.checksum
	}
	//fmt.Println(planetMap)
	fmt.Println("Task 6.1: ", sum)
	sanPath := make([]string, 0)
	planet = planetMap["SAN"]
	for planet.name != "COM" {
		sanPath = append(sanPath, planet.orbits)
		planet = planetMap[planet.orbits]
	}
	utils.ReverseSlice(sanPath)
	//fmt.Println(sanPath)

	youPath := make([]string, 0)
	planet = planetMap["YOU"]
	for planet.name != "COM" {
		youPath = append(youPath, planet.orbits)
		planet = planetMap[planet.orbits]
	}
	utils.ReverseSlice(youPath)
	//fmt.Println(youPath)

	commonPath := 0
	for true {
		if sanPath[commonPath] != youPath[commonPath] {
			break
		} else {
			commonPath++
		}
	}
	orbitalTrans := len(sanPath[commonPath:len(sanPath)]) + len(youPath[commonPath:len(youPath)])
	fmt.Println("Task 6.2: ", orbitalTrans)

}
