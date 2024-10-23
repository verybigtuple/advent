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

func IsNice2(input string) bool {
	pairs := make(map[Pair]int)
	var hasPairs, hasDouble bool
	var prev, prevPrev rune
	for i, cur := range input {
		if !hasPairs {
			if i == 0 {
				prev = cur
				continue
			}

			currentPair := Pair{prev, cur}
			if ind, ok := pairs[currentPair]; ok && (i-ind) > 1 {
				hasPairs = true
			} else {
				pairs[currentPair] = i
			}
		}

		if !hasDouble {
			if i == 0 {
				prev = cur
				continue
			}
			if i == 1 {
				prevPrev = prev
				prev = cur
				continue
			}

			hasDouble = prevPrev == cur
		}

		if hasPairs && hasDouble {
			return true
		}
		prevPrev = prev
		prev = cur
	}
	return false
}

func ReadFile() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	nices := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if IsNice(scanner.Text()) {
			nices++
		}
	}
	fmt.Printf("Read %v nice letters", nices)
	return nil
}

func main() {
	err := ReadFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(-1)
	}
}
