package main

import (
	"testing"

	"github.com/verybigtuple/advent/go2015-07/parser"
)

func TestComplexShema(t *testing.T) {
	m := map[string]interface{}{
		"x": parser.PureInput{123},
		"y": parser.PureInput{456},
		"d": parser.WiredBinary{parser.And, "x", "y"},
		"e": parser.WiredBinary{parser.Or, "x", "y"},
		"f": parser.Shift{parser.LShift, "x", 2},
		"g": parser.Shift{parser.RShift, "y", 2},
		"h": parser.Unary{parser.Not, "x"},
		"i": parser.Unary{parser.Not, "y"},
		"j": parser.WireInput{"x"},
	}

	want := map[string]uint16{
		"x": 123,
		"y": 456,
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"j": 123,
	}

	for wire := range m {
		t.Run(wire, func(t *testing.T) {
			got, err := CalcWire(m, wire)
			if err != nil {
				t.Errorf("Calc error %v", err)
			}
			w := want[wire]
			if got != w {
				t.Errorf("want: %d, got: {%d}", w, got)
			}
		})
	}

}
