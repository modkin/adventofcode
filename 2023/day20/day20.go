package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type module struct {
	modtype    string
	state      bool
	inputList  map[string]bool
	outputList []string
}

type pulse struct {
	src  string
	dest string
	high bool
}

func countBits(input map[string]bool) int {
	counter := 0
	for _, b := range input {
		if b {
			counter++
		}
	}
	return counter
}

func getPeriod(input map[string][]int) (int, bool) {
	valid := true
	if len(input) != 4 {
		valid = false
	}
	mult := 1
	for _, ints := range input {
		if len(ints) >= 2 {
			mult *= ints[len(ints)-1] - ints[len(ints)-2]
		} else {
			valid = false
		}
	}
	return mult, valid
}

func main() {
	file, err := os.Open("2023/day20/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	modules := make(map[string]module)
	for _, i2 := range lines {
		split := strings.Fields(i2)
		var modType string
		var modName string
		var outputList []string
		for _, s := range split[2:] {
			outputList = append(outputList, strings.Trim(s, ","))
		}
		if split[0][0] == 'b' {
			modType = split[0]
			modName = split[0]
		} else {
			modType = string(split[0][0])
			modName = split[0][1:]
		}
		newMod := module{
			modtype:    modType,
			outputList: outputList,
			inputList:  make(map[string]bool),
		}
		modules[modName] = newMod
	}
	for s, m := range modules {
		for _, s2 := range m.outputList {
			if _, ok := modules[s2]; !ok {
				fmt.Println(s2)
			} else {
				modules[s2].inputList[s] = false
			}
		}
	}

	period := make(map[string][]int)
	pressButton := func(iter int) ([2]int, bool) {
		pulses := [2]int{0, 0}
		rxPressed := false
		var pulseList []pulse
		pulses[0]++
		for _, s := range modules["broadcaster"].outputList {

			pulseList = append(pulseList, pulse{
				src:  "broadcast",
				dest: s,
				high: false,
			})
		}
		for len(pulseList) > 0 {
			for _, p := range pulseList {
				if p.high {
					pulses[1]++
				} else {
					pulses[0]++
				}
			}
			var newPulseList []pulse
			for _, p := range pulseList {
				targetMod := modules[p.dest]
				if targetMod.modtype == "%" {
					if !p.high {
						targetMod.state = !modules[p.dest].state
						for _, dest := range targetMod.outputList {
							newPulseList = append(newPulseList, pulse{p.dest, dest, targetMod.state})
						}
					}
				} else if targetMod.modtype == "&" {
					if p.dest == "cs" && p.high {
						period[p.src] = append(period[p.src], iter)
						number, valid := getPeriod(period)
						if valid {
							fmt.Println("Day 20.2:", number)
							rxPressed = true
						}

					}
					targetMod.inputList[p.src] = p.high
					for _, dest := range targetMod.outputList {

						newPulse := pulse{
							src:  p.dest,
							dest: dest,
							high: true,
						}
						if countBits(targetMod.inputList) == len(targetMod.inputList) {
							newPulse.high = false
						}
						newPulseList = append(newPulseList, newPulse)
					}
				}
				modules[p.dest] = targetMod
			}
			pulseList = newPulseList
		}
		return pulses, rxPressed
	}
	lowPulses := 0
	highPulses := 0
	for i := 0; true; i++ {
		pulses, rxPressed := pressButton(i)
		if rxPressed {
			break
		}
		lowPulses += pulses[0]
		highPulses += pulses[1]
		if i == 999 {
			fmt.Println("Day 20.1:", lowPulses*highPulses)
		}
	}

}
