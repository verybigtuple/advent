package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	want := Dem{1, 2, 100}

	input := fmt.Sprintf("%dx%dx%d", want.length, want.width, want.height)
	got, err := parseLine(input)
	if err != nil {
		t.Errorf("Error while parsing: %s", err.Error())
	}
	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestCalcWrapper(t *testing.T) {
	cases := []struct {
		name string
		dem  Dem
		want int
	}{
		{"zero", Dem{0, 0, 0}, 0},
		{"2x3x4", Dem{2, 3, 4}, 58},
		{"1x1x10", Dem{1, 1, 10}, 43},
		{"4x3x2", Dem{4, 3, 2}, 58},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := calcWrapperArea(c.dem)
			if c.want != got {
				t.Errorf("Want: %d, got: %d", c.want, got)
			}
		})
	}
}

func TestRibbonWrapper(t *testing.T) {
	cases := []struct {
		name string
		dem  Dem
		want int
	}{
		{"zero", Dem{0, 0, 0}, 0},
		{"2x3x4", Dem{2, 3, 4}, 34},
		{"1x1x10", Dem{1, 1, 10}, 14},
		{"4x3x2", Dem{4, 3, 2}, 34},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := calcRibbon(c.dem)
			if c.want != got {
				t.Errorf("Want: %d, got: %d", c.want, got)
			}
		})
	}
}
