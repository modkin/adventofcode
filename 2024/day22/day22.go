package main

import (
	"adventofcode/utils"
	"fmt"
	"time"
)

func secret(in int) int {

	const modulus = 16777216 // 2^24

	tmp1 := in * 64
	ret := tmp1 ^ in
	ret = ret % modulus
	tmp2 := ret / 32
	ret = tmp2 ^ ret
	ret = ret % modulus
	tmp3 := ret * 2048
	ret = tmp3 ^ ret
	ret = ret % modulus
	return ret

}

func find(list []int, test [4]int) (bool, int) {
	for i, i2 := range list {
		if i == len(list)-3 {
			return false, -1
		}
		if test[0] == i2 && test[1] == list[i+1] && test[2] == list[i+2] && test[3] == list[i+3] {
			return true, i + 3
		}
	}
	return false, -1
}

func main() {

	start := time.Now()
	lines := utils.ReadFileIntoLines("2024/day22/input")
	list := []int{}
	for _, line := range lines {
		list = append(list, utils.ToInt(line))
	}

	sum := 0
	for _, num := range list {
		for i := 0; i < 2000; i++ {
			num = secret(num)
		}
		sum += num
	}
	fmt.Println("Day 22.1:", sum)

	priceSeq := [][]int{}

	for _, num := range list {
		newPrices := []int{}
		for i := 0; i < 2000; i++ {
			newPrices = append(newPrices, num%10)
			num = secret(num)
		}
		priceSeq = append(priceSeq, newPrices)
	}
	diffSeq := [][]int{}
	for _, prices := range priceSeq {
		var newDiffSeq []int
		for i := 1; i < len(prices); i++ {
			newDiffSeq = append(newDiffSeq, prices[i]-prices[i-1])
		}
		diffSeq = append(diffSeq, newDiffSeq)
	}

	calcBananas := func(seq [4]int) int {
		sum := 0
		for i := 0; i < len(priceSeq); i++ {
			found, idx := find(diffSeq[i], seq)
			if found {
				tmp := priceSeq[i][idx+1]
				sum += tmp
			}
		}
		return sum
	}

	maxBananas := 0

	seqMap := make(map[[4]int]int)
	for customerId := 0; customerId < len(diffSeq); customerId++ {
		seq := [4]int{diffSeq[customerId][0], diffSeq[customerId][1], diffSeq[customerId][2], diffSeq[customerId][3]}
		for i := 3; i < len(diffSeq[customerId])-3; i++ {
			seq[3] = diffSeq[customerId][i]
			if _, ok := seqMap[seq]; !ok {
				seqMap[seq] = calcBananas(seq)
			}
			bananas := seqMap[seq]
			seq[0], seq[1], seq[2], seq[3] = seq[1], seq[2], seq[3], 0

			if bananas > maxBananas {
				maxBananas = bananas
			}
		}
	}

	fmt.Println("Day 22.2:", maxBananas)
	fmt.Println("Time taken:", time.Since(start))

}
