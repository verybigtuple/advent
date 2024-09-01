package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

type SoluttionFunc = func(io.Reader) int

func test(t *testing.T, name string, solution SoluttionFunc) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Empty", "", 0},
		{"One", "(", 1},
		{"One back", ")", -1},
		{"Forward back", "()", 0},
		{"Back Forward", ")(", 0},
		{"Positive", "(((", 3},
		{"Negative", ")))", -3},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v: %v", name, tt.name)
		t.Run(name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got := solution(reader)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, solFun SoluttionFunc) {

	var sb strings.Builder
	sb.Grow(1000)
	for i := 0; i < sb.Cap(); i++ {
		sb.WriteByte(leftBraket)
	}
	testString := sb.String()
	reader := strings.NewReader(testString)
	for i := 0; i < b.N; i++ {
		solFun(reader)
	}
}

func TestSolution0(t *testing.T) {
	test(t, "Solution0", solveFloor0)
}

func TestSolution1(t *testing.T) {
	test(t, "Solution1", solveFloor1)
}

func BenchmarkSoltion0(b *testing.B) {
	bench(b, solveFloor0)
}

func BenchmarkSoltion1(b *testing.B) {
	bench(b, solveFloor1)
}
