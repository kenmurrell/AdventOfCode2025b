package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func Part1(filename string) int {
	ranges := load(filename)
	count := 0
	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			if detect(i) {
				count += i
			}
		}
	}
	return count
}

func Part2(filename string) int {
	ranges := load(filename)
	count := 0
	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			if detectPart2(i) {
				count += i
			}
		}
	}
	return count
}

func detectPart2(number int) bool {
	s := strconv.Itoa(number)
	size := len(s)
	if size < 2 {
		return false
	}
	for blocklen := 1; blocklen <= size/2; blocklen++ {
		if size%blocklen != 0 {
			continue
		}
		block := s[:blocklen]
		queue := make(chan byte, size)
		for i := range size {
			queue <- block[i%blocklen]
		}
		close(queue)

		ok := true
		for i := range size {
			if s[i] != <-queue {
				ok = false
				break
			}
		}

		if ok {
			return true
		}
	}
	return false
}

func detect(number int) bool {
	runearr := []rune(strconv.Itoa(number))
	size := len(runearr)
	if size%2 != 0 {
		return false
	}
	queue := make(chan rune, size/2)
	for i := range size / 2 {
		queue <- runearr[i]
	}
	for i := size / 2; i < size; i++ {
		val1 := runearr[i]
		val2 := <-queue
		if val1 != val2 {
			return false
		}
	}
	return true
}

func load(filename string) []Range {
	file, _ := os.Open(filename)
	defer file.Close()
	var ranges []Range
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		textranges := strings.Split(text, ",")
		for _, r := range textranges {
			if len(r) == 0 {
				continue
			}
			n := strings.Split(r, "-")
			n1, _ := strconv.Atoi(n[0])
			n2, _ := strconv.Atoi(n[1])
			ranges = append(ranges, Range{n1, n2})
		}
	}
	return ranges
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER PART ONE: %d\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER PART TWO: %d\n", ans2)
}
