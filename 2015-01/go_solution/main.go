package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// or 40 in decimal for `(`
const leftBraket byte = 0x28

func solveFloor0(reader io.Reader) int {
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
func solveFloor1(reader io.Reader) int {
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

func solveBasement(reader io.Reader) int {
	br := bufio.NewReader(reader)

	i := 0
	c := 0
	for {
		ch, _, err := br.ReadRune()
		if err == io.EOF {
			break
		}

		c++
		if ch == '(' {
			i++
		} else {
			i--
		}

		if i == -1 {
			return c
		}
	}

	return -1
}

func run() error {
	f, err := os.Open("..\\input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	resultFloor := solveFloor0(f)
	f.Seek(0, io.SeekStart)
	resultBasement := solveBasement(f)
	fmt.Println("Floor in the end: ", resultFloor)
	fmt.Println("Basement at: ", resultBasement)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
