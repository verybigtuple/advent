package parser

type UnaryOperand = string
type BinaryOperand = string
type ShiftOperand = string

const arrow string = "->"

const (
	And    BinaryOperand = "AND"
	Or     BinaryOperand = "OR"
	LShift ShiftOperand  = "LSHIFT"
	RShift ShiftOperand  = "RSHIFT"
	Not    UnaryOperand  = "NOT"
	Empty  UnaryOperand  = ""
)

type PureInput struct {
	Input uint16
}

type WireInput struct {
	Input string
}

type Unary struct {
	Operand UnaryOperand
	Input   string
}

type WiredBinary struct {
	Operand BinaryOperand
	InputA  string
	InputB  string
}

type PureBinary struct {
	Operand BinaryOperand
	InputA  uint16
	InputB  string
}

type Shift struct {
	Operand ShiftOperand
	Input   string
	Param   byte
}
