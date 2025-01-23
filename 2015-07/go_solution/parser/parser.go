package parser

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
)

type ParsedLine struct {
	IntoWire  string
	Statement interface{} // PureInput, WireInput, Unary, WiredBinary, PureBinary, Shift
}


// Parser is a sctruct for parsing every line of input into a ParsedLine
type Parser struct {
	scanner   *bufio.Scanner
	line      int
	bufTokens [2]string
}

func New(src *bufio.Reader) *Parser {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	return &Parser{
		scanner: scanner,
		line:    0,
	}
}

func (p *Parser) NextLine() (*ParsedLine, error) {
	// we have to read 2 tokens in order to determine the possible type of the statement
	var err error
	p.bufTokens[0], err = p.readNextSrc()
	if errors.Is(err, ErrEOF) {
		return nil, err
	}
	p.bufTokens[1], err = p.readNextSrc()
	if errors.Is(err, ErrEOF) {
		return nil, &ParsingError{p.line, "unexpected end of file", err}
	}
	defer func() { p.line += 1 }()

	if p.bufTokens[0] == Not {
		return p.parseAsUnary()
	}

	switch p.bufTokens[1] {
	case arrow:
		if isNumeric(p.bufTokens[0]) {
			return p.parseAsPureInput()
		}
		return p.parseAsWireInput()
	case And, Or:
		if isNumeric(p.bufTokens[0]) {
			return p.parseAsPureBinary()
		}
		return p.parseAsWiredBinary()
	case LShift, RShift:
		return p.parseAsShift()
	default:
		return nil, &ParsingError{p.line, fmt.Sprintf("unexpected 2nd token %s", p.bufTokens[1]), nil}
	}
}

func (p *Parser) readNextSrc() (string, error) {
	if p.scanner.Scan() {
		return p.scanner.Text(), nil
	}
	return "", ErrEOF
}

func (p *Parser) getNextToken() (string, error) {
	if p.bufTokens[0] != "" {
		token := p.bufTokens[0]
		p.bufTokens[0] = p.bufTokens[1]
		p.bufTokens[1] = ""
		return token, nil
	}
	return p.readNextSrc()
}

func (p *Parser) expectArrowScan() (string, error) {
	arrowToken, err := p.getNextToken()
	if err != nil {
		return "", err
	}
	if errors.Is(err, ErrEOF) {
		return "", &ParsingError{p.line, "expected -> but got EOF", ErrEOF}
	}
	if arrowToken == arrow {
		return arrowToken, nil
	}
	return "", &ParsingError{p.line, fmt.Sprintf("expected ->, Got %s", arrowToken), nil}
}

func (p *Parser) expectIntScan() (uint16, error) {
	token, err := p.getNextToken()
	if errors.Is(err, ErrEOF) {
		return 0, &ParsingError{p.line, "expected integer but got EOF", ErrEOF}
	}
	input, err := strconv.ParseInt(token, 10, 16)
	if err != nil {
		return 0, &ParsingError{p.line, fmt.Sprintf("Cannot parse %v as integer", token), nil}
	}
	return uint16(input), nil
}

func (p *Parser) expectAlphaScan() (string, error) {
	token, err := p.getNextToken()
	if errors.Is(err, ErrEOF) {
		return "", &ParsingError{p.line, "expected non-numeric token but got EOF", ErrEOF}
	}
	if !isAlpha(token) {
		return "", &ParsingError{p.line, fmt.Sprintf("expected alpha token but got %s", token), nil}
	}
	return token, nil
}

func (p *Parser) parseAsUnary() (*ParsedLine, error) {
	// NOT x -> h
	_, err := p.getNextToken() // consume NOT as we have already checked it
	if err != nil {
		return nil, err
	}

	argWire, err := p.getNextToken() // x
	if err != nil {
		return nil, err
	}

	_, err = p.expectArrowScan() // ->
	if err != nil {
		return nil, err
	}

	intoWire, err := p.expectAlphaScan() // h
	if err != nil {
		return nil, err
	}

	parsedLine := ParsedLine{
		IntoWire:  intoWire,
		Statement: Unary{Not, argWire},
	}
	return &parsedLine, nil
}

func (p *Parser) parseAsPureInput() (*ParsedLine, error) {
	// 123 -> x
	input, err := p.expectIntScan() // 123
	if err != nil {
		return nil, err
	}

	_, err = p.expectArrowScan() // ->
	if err != nil {
		return nil, err
	}

	inputWire, err := p.expectAlphaScan() // x
	if err != nil {
		return nil, err
	}

	parsedLine := ParsedLine{
		IntoWire:  inputWire,
		Statement: PureInput{input},
	}
	return &parsedLine, nil
}

func (p *Parser) parseAsWireInput() (*ParsedLine, error) {
	// y -> x
	input, err := p.expectAlphaScan() // y
	if err != nil {
		return nil, err
	}

	_, err = p.expectArrowScan() // ->
	if err != nil {
		return nil, err
	}

	inputWire, err := p.expectAlphaScan() // x
	if err != nil {
		return nil, err
	}

	parsedLine := ParsedLine{
		IntoWire:  inputWire,
		Statement: WireInput{input},
	}
	return &parsedLine, nil
}

func (p *Parser) parseAsPureBinary() (*ParsedLine, error) {
	// 1 AND y -> d
	argA, err := p.expectIntScan() // 1
	if err != nil {
		return nil, err
	}

	opToken, err := p.getNextToken() // AND or OR
	if err != nil {
		return nil, err
	}
	op := And
	if opToken == Or {
		op = Or
	}

	argB, err := p.expectAlphaScan() // y
	if err != nil {
		return nil, err
	}

	_, err = p.expectArrowScan() // ->
	if err != nil {
		return nil, err
	}

	intoWire, err := p.expectAlphaScan() // d
	if err != nil {
		return nil, err
	}

	parsedLine := ParsedLine{
		IntoWire:  intoWire,
		Statement: PureBinary{op, argA, argB},
	}
	return &parsedLine, nil
}

func (p *Parser) parseAsWiredBinary() (*ParsedLine, error) {
	// x AND y -> d
	argA, err := p.expectAlphaScan() // x
	if err != nil {
		return nil, err
	}

	opToken, err := p.getNextToken() // AND or OR
	if err != nil {
		return nil, err
	}
	op := And
	if opToken == Or {
		op = Or
	}

	argB, err := p.expectAlphaScan() // y
	if err != nil {
		return nil, err
	}

	_, err = p.expectArrowScan() // ->
	if err != nil {
		return nil, err
	}

	intoWire, err := p.expectAlphaScan() // d
	if err != nil {
		return nil, err
	}

	parsedLine := ParsedLine{
		IntoWire:  intoWire,
		Statement: WiredBinary{op, argA, argB},
	}
	return &parsedLine, nil
}

func (p *Parser) parseAsShift() (*ParsedLine, error) {
	// x LSHIFT 2 -> f
	argA, err := p.expectAlphaScan() // x
	if err != nil {
		return nil, err
	}

	opToken, err := p.getNextToken() // LSHIFT
	if err != nil {
		return nil, err
	}
	op := LShift
	if opToken == RShift {
		op = RShift
	}

	shiftAmount, err := p.expectIntScan() // 2
	if err != nil {
		return nil, err
	}

	_, err = p.expectArrowScan() // ->
	if err != nil {
		return nil, err
	}
	intoWire, err := p.expectAlphaScan() //f
	if err != nil {
		return nil, err
	}

	parsedLine := ParsedLine{
		IntoWire:  intoWire,
		Statement: Shift{op, argA, byte(shiftAmount)},
	}
	return &parsedLine, nil
}

// isNumeric checks if a string is a number
func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// isAlpha checks if a string is alphabetic
func isAlpha(s string) bool {
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}
