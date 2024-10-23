package main

import "testing"

func TestIsNice(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"aaa", true},
		{"ugknbfddgicrmopn", true},
		{"jchzalrnumimnmhp", false}, // no doubles
		{"haegwjzuvuyypxyu", false}, // xy
		{"dvszwmarrgswjxmb", false}, // 1 vowel
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := IsNice(tc.input)
			if tc.want != got {
				t.Errorf("want: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestIsNice2(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"xxxx", true},
		{"aaa", false},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := IsNice2(tc.input)
			if tc.want != got {
				t.Errorf("want: %v, got: %v", tc.want, got)
			}
		})
	}
}
