package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1(filename string) int {
	moves := load(filename)
	i := 50
	count := 0
	for _, v := range moves {
		i = (i + v) % 100
		if i == 0 {
			count += 1
		}
	}
	return count
}

func Part2(filename string) int {
	moves := load(filename)
	i := 50
	count := 0
	for _, v := range moves {
		i0 := i
		i += v
		if i == 0 {
			count += 1
		}
		for i > 99 {
			i -= 100
			count += 1
		}
		for i < 0 {
			i += 100
			count += 1
		}
		if i0 == 0 && v < 0 {
			count--
		}
		fmt.Printf("rotate %d, goes to %d, count %d\n", v, i, count)
	}
	return count
}

func load(filename string) []int {
	file, _ := os.Open(filename)
	defer file.Close()
	var moves []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		dir := text[:1]
		num, _ := strconv.Atoi(text[1:])
		switch dir {
		case "L":
			moves = append(moves, num*-1)
		case "R":
			moves = append(moves, num)
		}
	}
	return moves
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER PART ONE: %d\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER PART TWO: %d\n", ans2)
}
