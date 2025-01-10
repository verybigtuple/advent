package parser

import (
	"bufio"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		input string
		want  *ParsedLine
	}{
		{"123 -> x", &ParsedLine{IntoWire: "x", Statement: PureInput{Input: 123}}},
		{"456 -> y", &ParsedLine{IntoWire: "y", Statement: PureInput{Input: 456}}},
		{"y -> x", &ParsedLine{IntoWire: "x", Statement: WireInput{Input: "y"}}},
		{"x AND y -> d", &ParsedLine{IntoWire: "d", Statement: WiredBinary{Operand: And, InputA: "x", InputB: "y"}}},
		{"x OR y -> e", &ParsedLine{IntoWire: "e", Statement: WiredBinary{Operand: Or, InputA: "x", InputB: "y"}}},
		{"1 AND y -> d", &ParsedLine{IntoWire: "d", Statement: PureBinary{Operand: And, InputA: 1, InputB: "y"}}},
		{"x LSHIFT 2 -> f", &ParsedLine{IntoWire: "f", Statement: Shift{Operand: LShift, Input: "x", Param: 2}}},
		{"y RSHIFT 2 -> g", &ParsedLine{IntoWire: "g", Statement: Shift{Operand: RShift, Input: "y", Param: 2}}},
		{"NOT x -> h", &ParsedLine{IntoWire: "h", Statement: Unary{Operand: Not, Input: "x"}}},
		{"NOT y -> i", &ParsedLine{IntoWire: "i", Statement: Unary{Operand: Not, Input: "y"}}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			p := New(bufio.NewReader(strings.NewReader(tc.input)))
			got, err := p.NextLine()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if *got != *tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
