package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// or 40 in decimal for `(`
const leftBraket byte = 0x28

func solveFirst0(reader io.Reader) int {
	br := bufio.NewReader(reader)

	i := 0
	for {
		ch, _, err := br.ReadRune()
		if err == io.EOF {
			break
		}
		if ch == '(' {
			i++
		} else {
			i--
		}
	}

	return i
}

// There is almost no difference between run and byte way.
func solveFirst1(reader io.Reader) int {
	br := bufio.NewReader(reader)

	i := 0
	for {
		ch, err := br.ReadByte()
		if err == io.EOF {
			break
		}
		if ch == leftBraket {
			i++
		} else {
			i--
		}
	}

	return i
}

func run() error {
	f, err := os.Open("..\\input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	result := solveFirst0(f)
	fmt.Println("Result: ", result)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
