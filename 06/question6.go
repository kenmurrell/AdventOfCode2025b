package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	A int
	B int
}

func Part1(filename string) int {
	values := load1(filename)
	sum := 0
	for i, _ := range values[0] {
		first, _ := strconv.Atoi(values[0][i])
		second, _ := strconv.Atoi(values[1][i])
		third, _ := strconv.Atoi(values[2][i])
		fourth, _ := strconv.Atoi(values[3][i])
		op := values[4][i]
		switch op {
		case "*":
			sum += first * second * third * fourth
		case "+":
			sum += first + second + third + fourth
		}
	}
	return sum
}

func load1(filename string) [][]string {
	file, _ := os.Open(filename)
	defer file.Close()
	var values [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		items := strings.Split(text, " ")
		var line []string
		for _, item := range items {
			if len(item) == 0 {
				continue
			}
			line = append(line, item)
		}
		values = append(values, line)

	}
	return values
}

func Part2(filename string) int64 {
	strArr := load2(filename)
	var total int64 = 0
	var numbers []int
	for col := len(strArr[0]) - 1; col >= 0; col-- {
		var warr []int
		for row := 0; row < len(strArr)-1; row++ {
			v := strArr[row][col]

			n, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			for i := range warr {
				warr[i] = warr[i] * 10
			}
			warr = append(warr, n)
		}
		if len(warr) == 0 {
			// prolly the gap between operations
			continue
		}
		newNumber := 0
		for i := range warr {
			newNumber += warr[i]
		}
		numbers = append(numbers, newNumber)

		v := strArr[len(strArr)-1][col]
		if v == "*" {
			subtotal := 1
			for i := range numbers {
				subtotal *= numbers[i]
			}
			total += int64(subtotal)
			numbers = numbers[:0]
			continue
		}
		if v == "+" {
			subtotal := 0
			for i := range numbers {
				subtotal += numbers[i]
			}
			total += int64(subtotal)
			numbers = numbers[:0]
			continue
		}
	}
	return total
}

func load2(filename string) [][]string {
	file, _ := os.Open(filename)
	defer file.Close()
	var values [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		items := strings.Split(text, "")
		var line []string
		for _, item := range items {
			line = append(line, item)
		}
		values = append(values, line)
	}
	return values
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER PART ONE: %d\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER PART TWO: %d\n", ans2)
}
