package main

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"", 1},
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			got, err := process(reader)
			if err != nil {
				t.Errorf("Error during process: %v", err)
			}
			if got != tc.want {
				t.Errorf("Want: %d, got: %d", tc.want, got)
			}
		})
	}
}
