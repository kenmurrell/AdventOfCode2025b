package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Empty struct{}

func Part1(filename string) int {
	devices := load(filename)
	visited := make(map[string]Empty)
	return traverse("you", devices, visited)
}

func traverse(current string, devices map[string][]string, visited map[string]Empty) int {
	if _, ok := visited[current]; ok {
		return 0
	}
	visited[current] = Empty{}
	if current == "out" {
		return 1
	}
	paths := 0
	conns := devices[current]
	for _, c := range conns {
		newVisited := copyMap(visited)
		paths += traverse(c, devices, newVisited)
	}
	return paths
}

func traverse2(current string, devices map[string][]string, visited map[string]Empty, hasDAC bool, hasFFT bool, cache map[string]int) int {
	if _, ok := visited[current]; ok {
		return 0
	}

	if current == "dac" {
		hasDAC = true
	}
	if current == "fft" {
		hasFFT = true
	}

	key := fmt.Sprintf("%s-%t-%t", current, hasDAC, hasFFT)
	if v, ok := cache[key]; ok {
		return v
	}

	visited[current] = Empty{}

	if current == "out" {
		if hasDAC && hasFFT {
			return 1
		}
		return 0
	}

	paths := 0
	for _, c := range devices[current] {
		newVisited := copyMap(visited)
		paths += traverse2(c, devices, newVisited, hasDAC, hasFFT, cache)
	}

	cache[key] = paths
	return paths
}

func copyMap[K comparable, V any](src map[K]V) map[K]V {
	dst := make(map[K]V, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func Part2(filename string) int {
	devices := load(filename)
	visited := make(map[string]Empty)
	cache := make(map[string]int)
	return traverse2("svr", devices, visited, false, false, cache)
}

func load(filename string) map[string][]string {
	file, _ := os.Open(filename)
	defer file.Close()
	devices := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		items := strings.Split(text, ": ")
		conns := strings.Split(items[1], " ")
		devices[items[0]] = conns
	}
	return devices
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER PART ONE: %d\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER PART TWO: %d\n", ans2)
}
