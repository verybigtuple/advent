package main


import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve("pqrstuv")
	}
}

func BenchmarkSolveWorkers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveWithWorkers("pqrstuv")
	}
}