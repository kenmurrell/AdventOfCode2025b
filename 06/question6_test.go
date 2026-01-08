package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ans := 4277556
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo(t *testing.T) {
	ans := 3263827
	rTest := Part2("test.txt")
	if rTest != int64(ans) {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}
