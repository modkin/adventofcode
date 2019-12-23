package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func cut(n int, deck []int) []int {
	if n > 0 {
		cut := deck[0:n]
		deck = append(deck[n:], cut...)
	} else if n < 0 {
		cut := deck[len(deck)+n:]
		deck = append(cut, deck[0:len(deck)+n]...)
	}
	return deck
}

func dealWithInc(inc int, deck []int) []int {
	newDeck := make([]int, len(deck))
	pos := 0
	for _, elem := range deck {
		newDeck[pos] = elem
		pos += inc
		pos = pos % len(deck)
	}
	return newDeck
}

func trackdealWithInc(inc int64, idx int64, len int64) int64 {
	prod1 := big.NewInt(inc)
	prod2 := big.NewInt(idx)
	lenB := big.NewInt(len)
	mult := big.NewInt(0)
	mult.Mul(prod1, prod2)
	mult.Mod(mult, lenB)
	return mult.Int64()

	//pos := int64(0)
	//for i := int64(0); i < idx; i++{
	//	pos += inc
	//	pos = pos%len
	//}
	//
	//return pos
	//remainderidx := idx%inc
	//remainderLength := len%inc
	//period := len/inc
	//if idx <= period {
	//	return idx * inc
	//}
	//
	//preOffset := int64(0)
	//var offset int64
	//for i := int64(0); i < idx/period; i++ {
	//	preOffset = offset
	//	if offset > remainderLength{
	//		offset = offset - remainderLength
	//	} else {
	//		offset = (offset + inc) - remainderLength
	//	}
	//	//offset = (preOffset + inc)%remainderLength - remainderLength
	//}
	//groupStart := remainderidx * inc
	//if groupStart == 0 {
	//	return (period -1) * inc + preOffset
	//}
	//if groupStart == 1 {
	//	if preOffset <= remainderLength {
	//		return period * inc + preOffset
	//	} else {
	//		return offset
	//	}
	//}
	//return (groupStart - inc) + offset

}

func trackReverse(idx int64, len int64) int64 {
	return len - idx - 1
}

func trackCut(cut int64, idx int64, len int64) int64 {
	if cut < 0 {
		cut = len + cut
	}
	if idx >= cut {
		return idx - cut
	} else {
		return idx + (len - cut)
	}
}

func main() {
	var deck []int
	for i := 0; i < 10007; i++ {
		deck = append(deck, i)
	}

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "cut" {
			deck = cut(utils.ToInt(line[1]), deck)
		} else if line[0] == "deal" && line[1] == "into" {
			deck = utils.ReverseIntSlice(deck)
		} else if line[0] == "deal" && line[1] == "with" {
			deck = dealWithInc(utils.ToInt(line[3]), deck)
		} else {
			fmt.Println("ERROR")
		}
	}
	for i, card := range deck {
		if card == 2019 {
			fmt.Println("Task 22.1: ", i)
		}
	}

	file, err = os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner = bufio.NewScanner(file)

	var funcs [][]int64
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "cut" {
			funcs = append(funcs, []int64{0, utils.ToInt64(line[1])})
		} else if line[0] == "deal" && line[1] == "into" {
			funcs = append(funcs, []int64{1})
		} else if line[0] == "deal" && line[1] == "with" {
			funcs = append(funcs, []int64{2, utils.ToInt64(line[3])})
		} else {
			fmt.Println("ERROR")
		}
	}

	pos := int64(2019)
	length := int64(10007)
	//pos = int64(2020)
	//length = int64(119315717514047)
	count := 101741582076661

	count = 1

	for i := 0; i < count; i++ {
		for _, inst := range funcs {
			switch inst[0] {
			case 0:
				pos = trackCut(inst[1], pos, length)
			case 1:
				pos = trackReverse(pos, length)
			case 2:
				pos = trackdealWithInc(inst[1], pos, length)
			}
		}
	}

	fmt.Println("Task 22.1: ", pos)

}
