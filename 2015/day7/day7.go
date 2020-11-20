package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("2015/day7/input.txt")
	if err != nil {
		panic(err)
	}

	wires := make(map[string]func() uint16)
	wireCache := make(map[string]uint16)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		inst := strings.Split(line, " ")
		if strings.Contains(line, "RSHIFT") {
			tmp := func() uint16 {
				if val, ok := wireCache[inst[0]]; ok {
					return val >> uint(utils.ToInt(inst[2]))
				} else {
					return wires[inst[0]]() >> uint(utils.ToInt(inst[2]))
				}
			}
			wires[inst[4]] = tmp
		} else if strings.Contains(line, "LSHIFT") {
			tmp := func() uint16 {
				return wires[inst[0]]() << uint(utils.ToInt(inst[2]))
			}
			wires[inst[4]] = tmp
		} else if strings.Contains(line, "OR") {
			tmp := func() uint16 {
				return wires[inst[0]]() | wires[inst[2]]()
			}
			wires[inst[4]] = tmp
		} else if strings.Contains(line, "AND") {
			val, err := strconv.Atoi(inst[0])
			var tmp func() uint16
			if err == nil {
				tmp = func() uint16 {
					return uint16(val) & wires[inst[2]]()
				}
			} else {
				tmp = func() uint16 {
					return wires[inst[0]]() & wires[inst[2]]()
				}
			}

			wires[inst[4]] = tmp
		} else if strings.Contains(line, "NOT") {
			tmp := func() uint16 {
				return ^wires[inst[1]]()
			}
			wires[inst[3]] = tmp
		} else {
			tmp := func() uint16 {
				val, err := strconv.Atoi(inst[0])
				if err == nil {
					return uint16(val)
				} else {
					return wires[inst[0]]()
				}
			}
			wires[inst[2]] = tmp
		}
	}

	//for key, val := range wires {
	//	fmt.Println(key, val())
	//}
	//fmt.Println(wires["z"]())
	fmt.Println("Task 7.1:", wires["a"]())
}
