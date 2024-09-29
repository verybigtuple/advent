package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math"
	"os"
	"sync"
)

const WORKER_COUNT int = 6

type Md5Hash = [16]byte

func CalcHash(input string, suffix uint) Md5Hash {
	str := fmt.Sprintf("%s%d", input, suffix)
	return md5.Sum([]byte(str))
}

// CheckPrefix5 checks if 5 first bytes in hex are zeros
func CheckPrefix5(h Md5Hash) bool {
	// To have 5 zeros in hex first byte 0 = 00, second 0 = 00  and third less 16 = 01, 02, ..., 0F
	return h[0] == 0x0 && h[1] == 0x0 && h[2] < 0x10
}

func CheckPrefix6(h Md5Hash) bool {
	return h[0] == 0x0 && h[1] == 0x0 && h[2] == 0x0
}

func Worker(wg *sync.WaitGroup, input string, hCheck func(Md5Hash) bool, suffix <-chan uint, result chan<- uint) {
	for s := range suffix {
		h := CalcHash(input, s)
		if hCheck(h) {
			result <- s
		}
		wg.Done()
	}
}

// SolveWithWorkers solving hashes using workers. Unfortunatelly, it is not effective.
// Md5 hash calc is too fast and it takes more time than sinle thread solution
func SolveWithWorkers(input string, hCheck func(Md5Hash) bool) (uint, error) {
	var wg sync.WaitGroup
	suffix := make(chan uint, WORKER_COUNT)
	result := make(chan uint, WORKER_COUNT)

	// Create workers
	for i := 0; i < WORKER_COUNT; i++ {
		go Worker(&wg, input, hCheck, suffix, result)
	}

	var i uint
	for {
		// Send numbers to workers
		for j := 0; j < WORKER_COUNT; j++ {
			if i == math.MaxUint {
				// hope I won't wait for this :)
				return 0, errors.New("cannot find the number: it is bigger than uint")
			}
			wg.Add(1)
			suffix <- i
			i++
		}

		// Wait until all workers are done in this cycle
		wg.Wait()

		// if we have result we should fetch it
		// theoreticaly we can have WORKER_COUNT results, so we should take the minimum
		if len(result) > 0 {
			var min uint = math.MaxUint
			for len(result) > 0 {
				r := <-result
				if r < min {
					min = r
				}
			}
			return min, nil
		}
	}
}

func Solve(input string, hCheck func(Md5Hash) bool) (uint, error) {
	for i := uint(0); i < math.MaxUint; i++ {
		h := CalcHash(input, i)
		if hCheck(h) {
			return i, nil
		}

	}

	return 0, errors.New("cannot find the number: it is bigger than uint")
}

func main() {
	str := "bgvyzdsv"
	s, err := SolveWithWorkers(str, CheckPrefix5)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(-1)
	}
	fmt.Printf("For \"%s\" solution %d\n", str, s)

	s, err = SolveWithWorkers(str, CheckPrefix6)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(-1)
	}
	fmt.Printf("For \"%s\" solution %d\n", str, s)
}
