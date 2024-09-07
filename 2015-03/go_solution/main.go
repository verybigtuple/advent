package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Point struct {
	x, y int
}

type PointSet = map[Point]struct{}

func move(v PointSet, current Point, dir rune) (Point, error) {
	switch dir {
	case '<':
		current.x -= 1
	case '>':
		current.x += 1
	case '^':
		current.y += 1
	case 'v':
		current.y -= 1
	default:
		return current, fmt.Errorf("unexpected symbol %v", dir)
	}

	v[current] = struct{}{}
	return current, nil
}

func process(reader io.Reader) (int, error) {
	bufReader := bufio.NewReader(reader)
	current := Point{0, 0}
	visited := PointSet{current: struct{}{}}

	for {
		r, _, err := bufReader.ReadRune()
		if err == io.EOF {
			break
		}
		current, err = move(visited, current, r)
		if err != nil {
			return 0, err
		}
	}
	return len(visited), nil
}

func runFile() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	houseCount, err := process(file)
	if err != nil {
		return err
	}
	fmt.Printf("Houses visited: %d", houseCount)
	return nil
}

func main() {
	err := runFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}
