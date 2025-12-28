package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Range struct {
	Start int
	End   int
}

func Part1(filename string) int {
	banks := load(filename)
	count := 0
	for _, b := range banks {
		n1, n2 := 0, 0
		n1_i := 0
		for i1 := 0; i1 < len(b)-1; i1++ {
			v, _ := strconv.Atoi(string(b[i1]))
			if v > n1 {
				n1 = v
				n1_i = i1
			}
		}
		for i2 := n1_i + 1; i2 < len(b); i2++ {
			v, _ := strconv.Atoi(string(b[i2]))
			if v > n2 {
				n2 = v
			}
		}
		count += (n1 * 10) + n2
	}
	return count
}

func Part2(filename string) int {
	banks := load(filename)
	count := 0
	for _, b := range banks {
		digits := make([]int, len(b))
		for i, c := range b {
			digits[i] = int(c - '0')
		}

		lastStart := 0
		var number int64
		for d := 11; d >= 0; d-- {
			i, n := findLargestInt(digits, lastStart, len(digits)-d)
			lastStart = i + 1
			number += int64(n) * int64(math.Pow10(d))
		}
		count += int(number)
	}
	return count
}

func findLargestInt(bank []int, start, end int) (int, int) {
	var n, n_i int = 0, 0
	for i := start; i < end; i++ {
		if bank[i] > n {
			n = bank[i]
			n_i = i
		}
	}
	return n_i, n
}

func load(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()
	var ranges []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		ranges = append(ranges, text)
	}
	return ranges
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER PART ONE: %d\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER PART TWO: %d\n", ans2)
}
