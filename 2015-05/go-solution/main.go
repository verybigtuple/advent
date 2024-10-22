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
