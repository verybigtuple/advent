package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

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

func CalcStatement(pLine parser.ParsedLine, calculatedWires map[string]uint16) (bool, error) {
	var err error
	switch s := pLine.Statement.(type) {
	case parser.PureInput:
		calculatedWires[pLine.IntoWire] = s.Input
		return true, nil
	case parser.WireInput:
		if _, ok := calculatedWires[s.Input]; ok {
			calculatedWires[pLine.IntoWire] = calculatedWires[s.Input]
		}
		return true, nil
	case parser.Unary:
		if _, ok := calculatedWires[s.Input]; ok {
			calculatedWires[pLine.IntoWire], err = CalcUnary(calculatedWires[s.Input], s.Operand)
			if err != nil {
				return false, err
			}
			return true, nil
		}

	case parser.PureBinary:
		if _, ok := calculatedWires[s.InputB]; ok {
			calculatedWires[pLine.IntoWire], err = CalcBinary(s.InputA, calculatedWires[s.InputB], s.Operand)
			if err != nil {
				return false, err
			}
			return true, nil
		}

	case parser.WiredBinary:
		_, okA := calculatedWires[s.InputA]
		_, okB := calculatedWires[s.InputB]
		if okA && okB {
			calculatedWires[pLine.IntoWire], err = CalcBinary(calculatedWires[s.InputA], calculatedWires[s.InputB], s.Operand)
			if err != nil {
				return false, err
			}
			return true, nil
		}

	case parser.Shift:
		if _, ok := calculatedWires[s.Input]; ok {
			calculatedWires[pLine.IntoWire], err = CalcShift(calculatedWires[s.Input], s.Param, s.Operand)
			if err != nil {
				return false, err
			}
			return true, nil
		}

	}
	return false, nil
}

func CalcWire(wires map[string]interface{}, wireName string) (uint16, error) {
	caclulatedWires := make(map[string]uint16)
	for {
		changed := false

		for wName, stat := range wires {
			if _, ok := caclulatedWires[wName]; ok {
				continue
			}

			var err error
			switch s := stat.(type) {
			case parser.PureInput:
				caclulatedWires[wName] = s.Input
				changed = true
			case parser.WireInput:
				if _, ok := caclulatedWires[s.Input]; ok {
					caclulatedWires[wName] = caclulatedWires[s.Input]
				}
				changed = true
			case parser.Unary:
				if _, ok := caclulatedWires[s.Input]; ok {
					caclulatedWires[wName], err = CalcUnary(caclulatedWires[s.Input], s.Operand)
					if err != nil {
						return 0, err
					}
					changed = true
				}

			case parser.PureBinary:
				if _, ok := caclulatedWires[s.InputB]; ok {
					caclulatedWires[wName], err = CalcBinary(s.InputA, caclulatedWires[s.InputB], s.Operand)
					if err != nil {
						return 0, err
					}
					changed = true
				}

			case parser.WiredBinary:
				_, okA := caclulatedWires[s.InputA]
				_, okB := caclulatedWires[s.InputB]
				if okA && okB {
					caclulatedWires[wName], err = CalcBinary(caclulatedWires[s.InputA], caclulatedWires[s.InputB], s.Operand)
					if err != nil {
						return 0, err
					}
					changed = true
				}

			case parser.Shift:
				if _, ok := caclulatedWires[s.Input]; ok {
					caclulatedWires[wName], err = CalcShift(caclulatedWires[s.Input], s.Param, s.Operand)
					if err != nil {
						return 0, err
					}
					changed = true
				}
			} // switch

			if _, ok := caclulatedWires[wireName]; ok {
				return caclulatedWires[wireName], nil
			}

		} // for	dict

		if !changed {
			return 0, fmt.Errorf("wire with the name %s cannot be resolved", wireName)
		}

	}
}

func run() error {
	f, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	p := parser.New(bufio.NewReader(f))
	wires := make(map[string]interface{})
	for {
		parsedLine, err := p.NextLine()
		if err != nil {
			if errors.Is(err, parser.EOF) {
				break
			}
			return err
		}
		wires[parsedLine.IntoWire] = parsedLine.Statement
	}
	result, err := CalcWire(wires, "a")
	if err != nil {
		return err
	}
	println(result)
	return nil

}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
