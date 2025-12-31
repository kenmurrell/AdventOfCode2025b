package main

import (
	"bufio"
	"fmt"
	"os"
)

func Part1(filename string) int {
	floor := load(filename)
	totalCount := 0
	for r := 0; r < len(floor); r++ {
		for c := 0; c < len(floor[r]); c++ {
			if floor[r][c] {
				count := countRollsAroundCoordinate(floor, r, c)
				if count < 4 {
					totalCount += 1
				}
			}
		}
	}
	return totalCount
}

func countRollsAroundCoordinate(floor [][]bool, r, c int) int {
	count := 0
	// check top
	count += checkCoordinate(floor, r-1, c)
	// check top right
	count += checkCoordinate(floor, r-1, c+1)
	// check right
	count += checkCoordinate(floor, r, c+1)
	// check bottom right
	count += checkCoordinate(floor, r+1, c+1)
	// check bottom
	count += checkCoordinate(floor, r+1, c)
	// check bottom left
	count += checkCoordinate(floor, r+1, c-1)
	// check left
	count += checkCoordinate(floor, r, c-1)
	// check top left
	count += checkCoordinate(floor, r-1, c-1)
	return count
}

func checkCoordinate(floor [][]bool, r, c int) int {
	if r < 0 || r >= len(floor) || c < 0 || c >= len(floor[r]) {
		return 0
	}
	if floor[r][c] {
		return 1
	} else {
		return 0
	}
}

func Part2(filename string) int {
	floor := load(filename)
	totalCount := 0
	for true {
		passCount := 0
		for r := 0; r < len(floor); r++ {
			for c := 0; c < len(floor[r]); c++ {
				if floor[r][c] {
					count := countRollsAroundCoordinate(floor, r, c)
					if count < 4 {
						floor[r][c] = false
						passCount += 1
					}
				}
			}
		}
		if passCount == 0 {
			break
		}
		totalCount += passCount
	}
	return totalCount
}

func load(filename string) [][]bool {
	file, _ := os.Open(filename)
	defer file.Close()
	var floor [][]bool
	scanner := bufio.NewScanner(file)
	for rowN := 0; scanner.Scan(); rowN++ {
		row := []rune(scanner.Text())
		rowF := make([]bool, len(row))
		for colN, v := range row {
			if v == '@' {
				rowF[colN] = true
			} else {
				rowF[colN] = false
			}
		}
		floor = append(floor, rowF)
	}
	return floor
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER PART ONE: %d\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER PART TWO: %d\n", ans2)
}
