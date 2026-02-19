package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	X int
	Y int
}

func Part1(filename string) float64 {
	tiles := load(filename)
	var max float64 = 0
	for i := 0; i < len(tiles); i++ {
		for j := i; j < len(tiles); j++ {
			width := math.Abs(float64(tiles[i].X-tiles[j].X)) + 1
			height := math.Abs(float64(tiles[i].Y-tiles[j].Y)) + 1
			area := width * height
			if area > max {
				max = area
			}
		}
	}
	return max
}

func Part2(filename string) float64 {
	return 0
}

func load(filename string) []Tile {
	file, _ := os.Open(filename)
	defer file.Close()
	var tiles []Tile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		items := strings.Split(text, ",")
		x, _ := strconv.Atoi(items[0])
		y, _ := strconv.Atoi(items[1])
		tiles = append(tiles, Tile{x, y})
	}
	return tiles
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER PART ONE: %f\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER PART TWO: %f\n", ans2)
}
