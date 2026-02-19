package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ans := 5
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo(t *testing.T) {
	ans := 2
	rTest := Part2("test2.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}
