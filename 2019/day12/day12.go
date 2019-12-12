package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Planet struct {
	pos [3]int
	vel [3]int
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func calcEnergy(planets [4]Planet) (energy int) {
	energy = 0
	for _, planet := range planets {
		kin, pot := 0, 0
		for i := 0; i <= 2; i++ {
			kin += utils.IntAbs(planet.pos[i])
			pot += utils.IntAbs(planet.vel[i])
		}
		energy += kin * pot
	}
	return
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	var planets [4]Planet
	scanner := bufio.NewScanner(file)
	planetPos := 0
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		var pos [3]int
		for i, planet := range coords {
			p, _ := strconv.Atoi(strings.TrimSuffix(strings.Split(planet, "=")[1], ">"))
			pos[i] = p
		}
		newPlanet := Planet{
			pos: pos,
			vel: [3]int{},
		}
		planets[planetPos] = newPlanet
		planetPos++
	}
	fmt.Println(planets)

	states := make(map[[4]Planet]int)
	//timesteps := 10 * 4686774924
	//for time := 0; time < timesteps; time++ {
	compPeriond := [3]int{0, 0, 0}

	///update velocity
	for component := 0; component < 3; component++ {
		time := 0
		for {
			states[planets] = time
			time++
			for idx, _ := range planets {
				for _, otherPlanet := range planets {
					if planets[idx] == otherPlanet {
						continue
					}

					if planets[idx].pos[component] > otherPlanet.pos[component] {
						planets[idx].vel[component]--
					} else if planets[idx].pos[component] < otherPlanet.pos[component] {
						planets[idx].vel[component]++
					}
				}
			}

			/// update position
			for idx, _ := range planets {
				planets[idx].pos[component] += planets[idx].vel[component]
			}

			if val, ok := states[planets]; ok {
				fmt.Println("Time: ", val)
				compPeriond[component] = time
				break
			}
			if time%10 == 0 {
				fmt.Println(time)
			}
		}
	}
	fmt.Println(compPeriond)
	fmt.Println(LCM(compPeriond[0], compPeriond[1], compPeriond[2]))

	fmt.Println(planets)

	fmt.Println("Task 12.1: ", calcEnergy(planets))
}
