package parser

import (
	"errors"
	"fmt"
)

var ErrEOF error = errors.New("EOF")

type ParsingError struct {
	Line    int
	Message string
	Err     error
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("%s at line %d", e.Message, e.Line)
}

func (e *ParsingError) Unwrap() error {
	return e.Err
}
