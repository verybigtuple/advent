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
