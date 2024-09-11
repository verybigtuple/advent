package main

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	testCases := []struct {
		input string
		wantSanta  int
		wantRobo int
	}{
		{"", 1, 1},
		{">", 2, 2},
		{"^>v<", 4, 3},
		{"^v^v^v^v^v", 2, 11},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			gotSanta, gotRobo, err := Process(reader)
			if err != nil {
				t.Errorf("Error during process: %v", err)
			}
			if gotSanta != tc.wantSanta {
				t.Errorf("Want Santa: %d, got Santa: %d", tc.wantSanta, gotSanta)
			}
			if gotRobo != tc.wantRobo {
				t.Errorf("Want Robo: %d, got Robo: %d", tc.wantRobo, gotRobo)
			}
		})
	}
}
