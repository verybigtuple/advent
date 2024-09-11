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

type innerPath struct {
	visited PointSet
}

func (ip innerPath) Len() int {
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

func (sp *SantaPath) Move(d rune) error {
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
		innerPath:    innerPath{visited: PointSet{startPoint: struct{}{}}},
	}
}

func (srp *SantaRobotPath) Move(d rune) error {
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

type PathMover interface {
	Len() int
	Move(rune) error
}

func CalcMovement(r rune, pm ...PathMover) error {
	for _, i := range pm {
		err := i.Move(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func Process(reader io.Reader) (int, int, error) {
	bufReader := bufio.NewReader(reader)
	santa := NewSantaPath()
	roboSants := NewSantaRobotPath()

	for {
		r, _, err := bufReader.ReadRune()
		if err == io.EOF {
			break
		}
		err = CalcMovement(r, santa, roboSants)
		if err != nil {
			return 0, 0, err
		}
	}
	return santa.Len(), roboSants.Len(), nil
}

func RunFile() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	houseCount1, houseCount2, err := Process(file)
	if err != nil {
		return err
	}
	fmt.Printf("Houses visited: %d, Robot visited %d", houseCount1, houseCount2)
	return nil
}

func main() {
	err := RunFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}
