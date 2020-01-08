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

func compact(shuffles [][]int64, cards int64) [][]int64 {
	var compacted [][]int64
	reverse := false
	for _, shuffle := range shuffles {
		if shuffle[0] == 1 {
			reverse = !reverse
			continue
		}
		if shuffle[0] != 1 && !reverse {
			compacted = append(compacted, shuffle)
		} else {
			if shuffle[0] == 0 {
				cut := (shuffle[1] + cards) % cards
				cut = cards - cut
				compacted = append(compacted, []int64{0, cut})
			} else if shuffle[0] == 2 {
				compacted = append(compacted, shuffle)
				compacted = append(compacted, []int64{0, cards + 1 - shuffle[1]})
			}
		}
	}
	if reverse {
		compacted = append(compacted, []int64{1})
	}
	shuffles = compacted

	compacted = make([][]int64, 0)
	cut := big.NewInt(0)
	cutInserted := false
	for _, shuffle := range shuffles {
		switch shuffle[0] {
		case 0:
			cut.Add(cut, big.NewInt(shuffle[1]))
			cut.Mod(cut, big.NewInt(cards))
		case 1:
			/// this might be wrong
			compacted = append(compacted, []int64{0, cut.Int64()})
			compacted = append(compacted, shuffle)
			cutInserted = true
		case 2:
			compacted = append(compacted, shuffle)
			cut.Mul(cut, big.NewInt(shuffle[1]))
			cut.Mod(cut, big.NewInt(cards))
		}
	}
	if !cutInserted {
		compacted = append(compacted, []int64{0, cut.Int64()})
	}

	shuffles = compacted

	compacted = make([][]int64, 0)
	increment := big.NewInt(1)
	dealWithIncInserted := false
	for _, shuffle := range shuffles {
		switch shuffle[0] {
		case 2:
			increment.Mul(increment, big.NewInt(shuffle[1]))
			increment.Mod(increment, big.NewInt(cards))
			//dealWithIncInserted = false
		default:
			if !dealWithIncInserted {
				compacted = append(compacted, []int64{2, increment.Int64()})
				dealWithIncInserted = true
			}
			compacted = append(compacted, shuffle)

		}
	}
	if !dealWithIncInserted {
		compacted = append(compacted, []int64{2, increment.Int64()})
	}

	return compacted
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

	pos := int64(2020)
	//length := int64(10007)
	//pos = int64(2020)
	length := int64(119315717514047)
	iterations := int64(101741582076661)

	//iterations = 1

	factor := compact(funcs, length)
	funcs = make([][]int64, 0)

	for iterLeft := length - iterations - 1; iterLeft != 0; iterLeft /= 2 {
		if iterLeft%2 == 1 {
			funcs = append(funcs, factor...)
			funcs = compact(funcs, length)
		}
		factor = append(factor, factor...)
		factor = compact(factor, length)
	}

	//for i := int64(0); i < iterations; i++ {
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
	//}

	fmt.Println("Task 22.2: ", pos)

}
