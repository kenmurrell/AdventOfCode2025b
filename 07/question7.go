package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Beam struct {
	Row int
	Col int
}

type empty struct{}

func Part1(filename string) int {
	manifold, b := load(filename)
	beams := make(map[Beam]empty)
	beams[b] = empty{}
	count := 0
	for true {
		if len(beams) == 0 {
			break
		}
		newBeams := make(map[Beam]empty)
		for b := range beams {
			delete(beams, b)

			// if at end, count it
			if b.Row == len(manifold)-1 {
				continue
			}
			if manifold[b.Row+1][b.Col] {
				count++
				if b.Col-1 >= 0 {
					newBeams[Beam{b.Row + 1, b.Col - 1}] = empty{}
				}
				if b.Col+1 < len(manifold[b.Row]) {
					newBeams[Beam{b.Row + 1, b.Col + 1}] = empty{}
				}
			} else {
				newBeams[Beam{b.Row + 1, b.Col}] = empty{}
			}
		}
		for b := range newBeams {
			beams[b] = empty{}
		}
	}

	return count
}

func Part2(filename string) int {
	manifold, b := load(filename)
	beams := make(map[Beam]int)
	beams[b] = 1
	count := 0
	for true {
		if len(beams) == 0 {
			break
		}
		newBeams := make(map[Beam]int)
		for b, v := range beams {
			delete(beams, b)

			// if at end, count it
			if b.Row == len(manifold)-1 {
				count += v
				continue
			}
			if manifold[b.Row+1][b.Col] {
				if b.Col-1 >= 0 {
					newBeams[Beam{b.Row + 1, b.Col - 1}] += v
				}
				if b.Col+1 < len(manifold[b.Row]) {
					newBeams[Beam{b.Row + 1, b.Col + 1}] += v
				}
			} else {
				newBeams[Beam{b.Row + 1, b.Col}] += v
			}
		}
		for b, v := range newBeams {
			beams[b] = v
		}
	}

	return count
}

func load(filename string) ([][]bool, Beam) {
	file, _ := os.Open(filename)
	defer file.Close()
	var manifold [][]bool
	var beam Beam
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowN := 0
		text := scanner.Text()
		items := strings.Split(text, "")
		row := make([]bool, len(items))
		for i, item := range items {
			switch item {
			case ".":
				row[i] = false
			case "^":
				row[i] = true
			case "S":
				beam = Beam{rowN, i}
			}
		}
		manifold = append(manifold, row)
		rowN++
	}
	return manifold, beam
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER PART ONE: %d\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER PART TWO: %d\n", ans2)
}

// 17815 too high
