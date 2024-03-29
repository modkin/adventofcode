package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toIntList(input int) []int {
	output := make([]int, 0)
	intString := strconv.Itoa(input)
	for _, s := range strings.Split(intString, "") {
		output = append(output, utils.ToInt(s))
	}
	return output
}

func main() {
	file, err := os.Open("2021/day24/testinput")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	registers := make(map[string]int)
	registers["w"] = 0
	registers["x"] = 0
	registers["y"] = 0
	registers["z"] = 0
	input := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9}

	alu := func(inst string, ops []string) {
		var op1 int
		if len(ops) == 2 {
			ret, err := strconv.Atoi(strings.TrimSpace(ops[1]))
			if err != nil {
				op1 = registers[ops[1]]
			} else {
				op1 = ret
			}
		}
		switch inst {
		case "inp":
			fmt.Println("---------------------")
			registers[ops[0]] = input[0]
			input = input[1:]
		case "mul":
			registers[ops[0]] *= op1
		case "add":
			registers[ops[0]] += op1
		case "div":
			registers[ops[0]] /= op1
		case "mod":
			registers[ops[0]] %= op1
		case "eql":
			if registers[ops[0]] == registers[ops[1]] {
				registers[ops[0]] = 1
			} else {
				registers[ops[0]] = 0
			}
		}
	}

	inst := make([][]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		inst = append(inst, line)

	}
	check := func() bool {
		//registers["w"] = 0
		//registers["x"] = 0
		//registers["y"] = 0
		//registers["z"] = 0
		for _, in := range inst {
			alu(in[0], in[1:])
			if in[0] == "add" && in[1] == "x" {
				fmt.Println(in)
				fmt.Println(registers)
			}
		}
		if len(input) != 0 {
			fmt.Println("EROOR")
		}
		if registers["z"] == 0 {
			return true
		} else {
			return false
		}
	}

	i := 78896919699991
	i = 59996912981939 // part 1
	i = 17241911811915 //part 2
	registers["z"] = 0

	//for k := 0; k < 9; k++ {
	input = toIntList(i)
	fmt.Println(input)
	check()
	fmt.Println(registers)
	//i -= 10000000000000
	//}

}
