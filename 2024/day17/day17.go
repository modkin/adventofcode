package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func getComb(reg map[string]int, operand int) int {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return reg["A"]
	} else if operand == 5 {
		return reg["B"]
	} else if operand == 6 {
		return reg["C"]
	}
	fmt.Println("ERROR")
	return math.MaxInt
}

func perfOp(regs map[string]int, prog []int, opcode, operand int, output []int) (int, []int) {

	if opcode == 0 {
		regs["A"] = int(float64(regs["A"]) / math.Pow(float64(2), float64(getComb(regs, operand))))
	}

	if opcode == 1 {
		regs["B"] = regs["B"] ^ operand
	}

	if opcode == 2 {
		regs["B"] = getComb(regs, operand) % 8
	}

	if opcode == 3 {
		if regs["A"] != 0 {
			return int(operand), output
		}
	}

	if opcode == 4 {
		regs["B"] = regs["B"] ^ regs["C"]
	}

	if opcode == 5 {
		//fmt.Print(getComb(regs, operand)%8, ",")
		output = append(output, getComb(regs, operand)%8)
	}

	if opcode == 6 {
		regs["B"] = int(float64(regs["A"]) / math.Pow(float64(2), float64(getComb(regs, operand))))
	}

	if opcode == 7 {
		regs["C"] = int(float64(regs["A"]) / math.Pow(float64(2), float64(getComb(regs, operand))))
	}
	return -1, output
}

func runProgramm(regs map[string]int, programm []int) []int {
	output := []int{}
	newPos := 0
	for ptr := int(0); ptr < int(len(programm)); {
		newPos, output = perfOp(regs, programm, programm[ptr], programm[ptr+1], output)
		if newPos != -1 {
			ptr = int(newPos)
		} else {
			ptr += 2
		}
	}
	return output
}

func splitIntoChunks(s string, chunkSize int) []string {
	var chunks []string
	length := len(s)

	for i := length; i > 0; i -= chunkSize {
		start := i - chunkSize
		if start < 0 {
			start = 0
		}
		chunks = append([]string{s[start:i]}, chunks...) // Prepend to maintain order
	}

	return chunks
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day17/input")

	reg := regexp.MustCompile(`Register (.): (\d*)`)

	regs := make(map[string]int)
	programm := []int{}

	readsregs := true
	for _, line := range lines {
		if line == "" {
			readsregs = false
			continue
		}
		if readsregs {
			r := reg.FindStringSubmatch(line)
			regs[r[1]] = utils.ToInt(r[2])
		} else {
			split := strings.Split(line, " ")
			split2 := strings.Split(split[1], ",")
			for _, s := range split2 {
				programm = append(programm, utils.ToInt(s))
			}

		}
	}
	fmt.Println(regs)
	fmt.Println(programm)
	//
	//regs["A"] = 0
	//regs["B"] = 2024
	//regs["C"] = 43690
	//programm = []int{4, 0}

	var output []int
	//var newPos int

	//outer:
	outToInt := make(map[int][]int)
	//int(221190396313600)
	//110010010010101111100000000000000
	//110010010010101111100000100000000000
	//110010010010101111100000100010000000000
	//110010010010101111100000100010101000000000
	//110010010010101111100000100010101000000000000000
	//110101110010101101000000000
	//110101110010101101110000|000000
	//110|101|110|010|101|101|110|000|000000
	// 0 | 3 | 5 | 5 | 5 | 1 | 3 |
	binaryStr := "110101110010101101110000000000"
	binaryStr = strings.ReplaceAll("110|101|110|010|100|101|111|100|000|000000000000000000000", "|", "")
	stopuntin := strings.ReplaceAll("110|101|110|010|100|101|111|101|000|000000000000000000000", "|", "")

	binaryStr = strings.ReplaceAll("110|101|110|010|010|101|111|100|000|000|000|000|000|000|000|000", "|", "")
	stopuntin = strings.ReplaceAll("110|101|110|010|010|101|111|100|001|000|000|000|000|000|000|000", "|", "")
	intValue, err := strconv.ParseInt(binaryStr, 2, 64) // Base 2, 64-bit integer
	stopValue, _ := strconv.ParseInt(stopuntin, 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

outer:
	for i := int(intValue); i < int(stopValue); i += 1 {

		regs["A"] = i
		regs["B"] = 0
		regs["C"] = 0
		output = make([]int, 0)
		//regs["A"] = 117440
		output = runProgramm(regs, programm)
		checkAll := true
		for i2 := 0; i2 < len(output)-3; i2++ {
			if output[len(output)-1-i2] != programm[len(programm)-1-i2] {
				checkAll = false
			}
		}

		if checkAll {
			fmt.Println("found")
			fmt.Println(output)
			fmt.Println(strconv.FormatInt(int64(i), 2))
			split := splitIntoChunks(strconv.FormatInt(int64(i), 2), 3)
			fmt.Println(split)
			fmt.Println(len(split), "/", len(programm))
			fmt.Println()
			//break outer
		}

		//fmt.Println(i, output)
		if i < 8 {
			outToInt[output[0]] = append(outToInt[output[0]], i)
		}
		//fmt.Println(output)
		if len(programm) != len(output) {
			continue outer
		}
		for i2, u := range programm {
			if output[i2] != u {
				continue outer
			}
		}

		fmt.Println("Day 17.2:", i)
		break
	}

	fmt.Println(outToInt)
	fmt.Println(regs)
	fmt.Println(output)
}
