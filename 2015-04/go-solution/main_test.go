package main


import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve("bgvyzdsv", CheckPrefix6)
	}
}

func BenchmarkSolveWorkers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveWithWorkers("bgvyzdsv", CheckPrefix6)
	}
}