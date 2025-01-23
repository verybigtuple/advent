package main

import (
	"testing"

	"github.com/verybigtuple/advent/go2015-07/parser"
)

func TestComplexShema(t *testing.T) {
	testCases := []struct {
		parsedLine *parser.ParsedLine
		want       uint16
	}{
		{&parser.ParsedLine{"x", parser.PureInput{123}}, 123},
		{&parser.ParsedLine{"y", parser.PureInput{456}}, 456},
		{&parser.ParsedLine{"d", parser.WiredBinary{parser.And, "x", "y"}}, 72},
		{&parser.ParsedLine{"e", parser.WiredBinary{parser.Or, "x", "y"}}, 507},
		{&parser.ParsedLine{"f", parser.Shift{parser.LShift, "x", 2}}, 492},
		{&parser.ParsedLine{"g", parser.Shift{parser.RShift, "y", 2}}, 114},
		{&parser.ParsedLine{"h", parser.Unary{parser.Not, "x"}}, 65412},
		{&parser.ParsedLine{"i", parser.Unary{parser.Not, "y"}}, 65079},
		{&parser.ParsedLine{"j", parser.WireInput{"x"}}, 123},
	}

	wires := make([]*parser.ParsedLine, 0, len(testCases))
	for _, tc := range testCases {
		wires = append(wires, tc.parsedLine)
	}

	for _, tc := range testCases {
		t.Run(tc.parsedLine.IntoWire, func(t *testing.T) {
			got, err := CalcWire(wires, tc.parsedLine.IntoWire)
			if err != nil {
				t.Errorf("Calc error %v", err)
			}
			w := tc.want
			if got != w {
				t.Errorf("want: %d, got: {%d}", w, got)
			}
		})
	}

}
