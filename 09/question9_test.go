package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	var ans float64 = 50
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %f; want %f", rTest, ans)
	}
}

func TestPartTwo(t *testing.T) {
	ans := 40.0
	rTest := Part2("test.txt")
	if rTest != ans {
		t.Errorf("Got %f; want %f", rTest, ans)
	}
}
