package codingchallenges

import (
	"fmt"
	"testing"
)

/* The first coding interview Om Imre's had, asked to Dzuizz
Input: 	n=12
Output: 1 * 3 4 * 6 * 8 9 * 11 * 12

Input: 	n=20
Output: 1 * 3 4 * 6 * 8 9 * 11 * 13 14 * 16 * 18 19 *


Pattern: Bintang kelipatan 5, selang seling.
*/

func PrintStarsWithPattern(num int) {
	for i := 1; i <= num; i++ {
		if i%5 == 2 || i%5 == 0 {
			fmt.Printf("* ")
		} else {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Printf("\n")
}

func TestPrintStarsWithPattern(t *testing.T) {
	PrintStarsWithPattern(12)
	/*
		=== RUN   TestPrintStarsWithPattern
		1 * 3 4 * 6 * 8 9 * 11 *
		--- PASS: TestPrintStarsWithPattern (0.00s)
		PASS
	*/

	PrintStarsWithPattern(20)
	/*
			=== RUN   TestPrintStarsWithPattern
		1 * 3 4 * 6 * 8 9 * 11 * 13 14 * 16 * 18 19 *
		--- PASS: TestPrintStarsWithPattern (0.00s)
		PASS
		ok      coding-challenges       0.389s
	*/
}
