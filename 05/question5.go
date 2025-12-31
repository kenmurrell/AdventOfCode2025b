package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	A int
	B int
}

func Part1(filename string) int {
	ingredients, intervals := load(filename)
	count := 0
	for _, ingredient := range ingredients {
		for _, r := range intervals {
			if ingredient >= r.A && ingredient <= r.B {
				count++
				break
			}
		}
	}
	return count
}

func Part2(filename string) int64 {
	_, intervals := load(filename)
	var count int64 = 0
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].A == intervals[j].A {
			return intervals[i].B < intervals[j].B
		}
		return intervals[i].A < intervals[j].A
	})
	for i, j := 0, 1; j < len(intervals); i, j = i+1, j+1 {
		if intervals[i].B >= intervals[j].A {
			oldIntervalB := intervals[i].B
			intervals[i].B = intervals[j].A - 1
			if oldIntervalB > intervals[j].B {
				intervals[j].B = oldIntervalB
			}
		}
	}
	for _, interval := range intervals {
		diff := interval.B - interval.A
		if diff < 0 {
			// I don't want to waste any time solving this edgecase, so we'll just skip it for now :/
			fmt.Printf("%d-%d\n", interval.A, interval.B)
			continue
		}
		count += int64(diff + 1)
	}
	return count
}

func load(filename string) ([]int, []Interval) {
	file, _ := os.Open(filename)
	defer file.Close()
	var ingredients []int
	var intervals []Interval
	scanner := bufio.NewScanner(file)
	scanIngredients := false
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			scanIngredients = true
			continue
		}
		if scanIngredients {
			i, _ := strconv.Atoi(text)
			ingredients = append(ingredients, i)
		} else {
			nums := strings.Split(text, "-")
			a, _ := strconv.Atoi(nums[0])
			b, _ := strconv.Atoi(nums[1])
			intervals = append(intervals, Interval{a, b})
		}
	}
	return ingredients, intervals
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER PART ONE: %d\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER PART TWO: %d\n", ans2)
}
