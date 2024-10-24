package main

import "testing"

type TestCase struct {
	input string
	want bool
}


func test(t *testing.T, testCases []TestCase, testFunc func(string) bool) {
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := testFunc(tc.input)
			if tc.want != got {
				t.Errorf("want: %v, got: %v", tc.want, got)
			}
		})
	}
}


func TestIsNice(t *testing.T) {
	tc := []TestCase{
		{"aaa", true},
		{"ugknbfddgicrmopn", true},
		{"jchzalrnumimnmhp", false}, // no doubles
		{"haegwjzuvuyypxyu", false}, // xy
		{"dvszwmarrgswjxmb", false}, // 1 vowel
	}
	test(t, tc, IsNice)
}

func TestIsNice2(t *testing.T) {
	tc := []TestCase{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"xxxx", true},
		{"aaa", false},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	test(t, tc, IsNice2)
}
