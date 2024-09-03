package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dem struct {
	length int
	width  int
	height int
}

func parseLine(line string) (Dem, error) {
	var d Dem
	strReader := strings.NewReader(line)
	_, err := fmt.Fscanf(strReader, "%dx%dx%d", &d.length, &d.width, &d.height)
	if err != nil {
		return d, err
	}
	return d, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcWrapperArea(d Dem) int {
	a, b, c := d.length*d.width, d.width*d.height, d.height*d.length
	m := min(min(a, b), c)
	return 2*a + 2*b + 2*c + m
}

func calcRibbon(d Dem) int {
	m := max(max(d.height, d.length), d.width)
	wrap := 2*d.height + 2*d.length + 2*d.width - 2*m
	bow := d.height * d.length * d.width
	return wrap + bow
}

func run() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("file cannot be open %w", err)
	}
	defer file.Close()

	var sumWrap, sumRibbon int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		d, err := parseLine(line)
		if err != nil {
			return fmt.Errorf("cannot parse line %v: %w", line, err)
		}
		sumWrap += calcWrapperArea(d)
		sumRibbon += calcRibbon(d)
	}
	fmt.Printf("Wrapper: %d\n", sumWrap)
	fmt.Printf("Ribbon: %d\n", sumRibbon)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
