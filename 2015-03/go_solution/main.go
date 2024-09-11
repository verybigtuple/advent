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

func shift(current Point, direction rune) (Point, error) {
	switch direction {
	case '<':
		current.x -= 1
	case '>':
		current.x += 1
	case '^':
		current.y += 1
	case 'v':
		current.y -= 1
	default:
		return current, fmt.Errorf("unexpected symbol %v", direction)
	}
	return current, nil
}

type PathMover interface {
	len() int
	move(rune) error
}

type innerPath struct {
	visited PointSet
}

func (ip innerPath) len() int {
	return len(ip.visited)
}

type SantaPath struct {
	current Point
	innerPath
}

func NewSantaPath() *SantaPath {
	startPoint := Point{0, 0}
	return &SantaPath{
		current:   startPoint,
		innerPath: innerPath{visited: PointSet{startPoint: struct{}{}}},
	}
}

func (sp *SantaPath) move(d rune) error {
	var err error
	sp.current, err = shift(sp.current, d)
	if err != nil {
		return err
	}
	sp.visited[sp.current] = struct{}{}
	return nil
}

type SantaRobotPath struct {
	currentSanta Point
	currentRobot Point
	santaMoves   bool
	innerPath
}

func NewSantaRobotPath() *SantaRobotPath {
	startPoint := Point{0, 0}
	return &SantaRobotPath{
		currentSanta: startPoint,
		currentRobot: startPoint,
		santaMoves:   true,
		innerPath: innerPath{visited: PointSet{startPoint: struct{}{}}},
	}
}

func (srp *SantaRobotPath) move(d rune) error {
	var err error
	var p Point
	if srp.santaMoves {
		p, err = shift(srp.currentSanta, d)
		srp.currentSanta = p
	} else {
		p, err = shift(srp.currentRobot, d)
		srp.currentRobot = p
	}
	if err != nil {
		return err
	}
	srp.santaMoves = !srp.santaMoves

	srp.visited[p] = struct{}{}
	return nil
}


func process(reader io.Reader) (int, error) {
	bufReader := bufio.NewReader(reader)
	santa := NewSantaPath()

	for {
		r, _, err := bufReader.ReadRune()
		if err == io.EOF {
			break
		}
		err = santa.move(r)
		if err != nil {
			return 0, err
		}
	}
	return santa.len(), nil
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
