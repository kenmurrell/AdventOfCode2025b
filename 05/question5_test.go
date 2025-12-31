package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ans := 3
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo(t *testing.T) {
	ans := 14
	rTest := Part2("test.txt")
	if rTest != int64(ans) {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}
