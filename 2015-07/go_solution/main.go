package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/verybigtuple/advent/go2015-07/parser"
)

func CalcUnary(input uint16, operand parser.UnaryOperand) (uint16, error) {
	switch operand {
	case parser.Not:
		return ^input, nil
	}
	return 0, errors.ErrUnsupported
}

func CalcBinary(inputA, inputB uint16, operand parser.BinaryOperand) (uint16, error) {
	switch operand {
	case parser.And:
		return inputA & inputB, nil
	case parser.Or:
		return inputA | inputB, nil
	}
	return 0, errors.ErrUnsupported
}

func CalcShift(inputA uint16, param byte, operand parser.ShiftOperand) (uint16, error) {
	switch operand {
	case parser.LShift:
		return inputA << param, nil
	case parser.RShift:
		return inputA >> param, nil
	}
	return 0, errors.ErrUnsupported
}

// CalcStatement calculates the current wire if possible. If the wire cannot be calculated, returns false as the second return value.
// It does not change the calculatedWires map.
func CalcStatement(pLine *parser.ParsedLine, calculatedWires map[string]uint16) (uint16, bool, error) {
	switch s := pLine.Statement.(type) {
	case parser.PureInput:
		return s.Input, true, nil
	case parser.WireInput:
		if input, ok := calculatedWires[s.Input]; ok {
			return input, true, nil
		}
	case parser.Unary:
		if input, ok := calculatedWires[s.Input]; ok {
			unary, err := CalcUnary(input, s.Operand)
			if err != nil {
				return 0, false, err
			}
			return unary, true, nil
		}
	case parser.PureBinary:
		if inputB, ok := calculatedWires[s.InputB]; ok {
			binary, err := CalcBinary(s.InputA, inputB, s.Operand)
			if err != nil {
				return 0, false, err
			}
			return binary, true, nil
		}
	case parser.WiredBinary:
		inputA, okA := calculatedWires[s.InputA]
		inputB, okB := calculatedWires[s.InputB]
		if okA && okB {
			binary, err := CalcBinary(inputA, inputB, s.Operand)
			if err != nil {
				return 0, false, err
			}
			return binary, true, nil
		}
	case parser.Shift:
		if input, ok := calculatedWires[s.Input]; ok {
			shift, err := CalcShift(input, s.Param, s.Operand)
			if err != nil {
				return 0, false, err
			}
			return shift, true, nil
		}

	}
	return 0, false, nil
}

func CalcWire(wires []*parser.ParsedLine, wireName string) (uint16, error) {
	caclulatedWires := make(map[string]uint16)
	for {
		changed := false

		for _, parsedLine := range wires {
			if _, ok := caclulatedWires[parsedLine.IntoWire]; ok {
				continue
			}

			value, isCalc, err := CalcStatement(parsedLine, caclulatedWires)
			if err != nil {
				return 0, err
			}
			if isCalc {
				caclulatedWires[parsedLine.IntoWire] = value
			}
			changed = changed || isCalc

			if calcWire, ok := caclulatedWires[wireName]; ok {
				return calcWire, nil
			}

		}

		if !changed {
			return 0, fmt.Errorf("wire with the name %s cannot be resolved", wireName)
		}
	}
}

func readAllWires(f *os.File) ([]*parser.ParsedLine, error) {
	p := parser.New(bufio.NewReader(f))
	wires := make([]*parser.ParsedLine, 0)
	for {
		parsedLine, err := p.NextLine()
		if err != nil {
			if errors.Is(err, parser.ErrEOF) {
				break
			}
			return nil, err
		}
		wires = append(wires, parsedLine)
	}
	return wires, nil
}

func run() error {
	f, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	wires, err := readAllWires(f)
	if err != nil {
		return err
	}

	// calculate the firtst part of the problem
	result1, err := CalcWire(wires, "a")
	if err != nil {
		return err
	}
	fmt.Printf("Part1 answer: %d\n", result1)

	// Transform as per the problem statement
	bWireIdx := slices.IndexFunc(wires, func(w *parser.ParsedLine) bool { return w.IntoWire == "b" })
	if bWireIdx == -1 {
		return errors.New("wire b not found")
	}
	wires[bWireIdx].Statement = parser.PureInput{Input: result1}

	// calculate the second part of the problem
	result2, err := CalcWire(wires, "a")
	if err != nil {
		return err
	}
	fmt.Printf("Part2 answer: %d\n", result2)

	return nil

}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
