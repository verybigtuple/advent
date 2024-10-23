package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsNice(input string) bool {
	var prev rune
	var vowels, pairs int
	for _, cur := range input {
		switch cur {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'b', 'd', 'q', 'y':
			if cur-prev == 1 {
				return false
			}
		}
		if cur == prev {
			pairs++
		}
		prev = cur
	}
	return vowels >= 3 && pairs >= 1
}

type Pair struct {
	R1 rune
	R2 rune
}

// RuneBuf saves history of 2 last runes that are necessary for the second case
type RuneBuf [2]rune

func (rb *RuneBuf) Push(r rune) {
	rb[0], rb[1] = r, rb[0]
}

func IsNice2(input string) bool {
	// int in this map is index of current rune
	// using this index we should determine if pairs overlap
	pairs := make(map[Pair]int)
	var buf RuneBuf
	var hasPairs, hasDouble bool
	for i, cur := range input {
		if !hasPairs && buf[0] != 0 {
			currentPair := Pair{buf[0], cur}
			ind, ok := pairs[currentPair]
			// in case `aaa` aa overlaps
			// so the pair will be written in the map as (a, a): 1
			// the next pair (a, a):2, so 2-1 = 1. Pairs overlap
			if ok && (i-ind) > 1 {
				hasPairs = true
			}
			// cannot set else here: it might be ok but (i-ind) == 1, in this case I do nothing
			if !ok {
				pairs[currentPair] = i
			}
		}

		if !hasDouble && buf[1] != 0 {
			hasDouble = buf[1] == cur
		}

		if hasPairs && hasDouble {
			return true
		}
		buf.Push(cur)
	}
	return false
}

func ReadFile() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	var nices1, nices2 int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if IsNice(line) {
			nices1++
		}
		if IsNice2(line) {
			nices2++
		}
	}
	fmt.Printf("Read %v nice letters with 1st criteria \n", nices1)
	fmt.Printf("Read %v nice letters with 2nd criteria \n", nices2)
	return nil
}

func main() {
	err := ReadFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(-1)
	}
}
