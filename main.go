// main recursively solves the puzzle described in the README. By default it
// solves the 8-digit version. Two solutions exist for the 6-digit puzzle; other
// digit counts result in an infinite recursion.
//
// TODO: detect the recusive loop, label the numbers involved as invalid, and
// continue to solve. Do solutions exist?
package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	digits = 8
)

func main() {
	flag.Parse()

	switch flag.NArg() {
	case 0: // do nothing
	case 1:
		arg := flag.Arg(0)
		var err error
		if digits, err = strconv.Atoi(arg); err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintf("I don't know how to handle %d args.", flag.NArg()))
	}

	if digits < 1 || digits > 9 {
		panic("Arg must be in range 1-9.")
	}

	fmt.Printf("Running puzzle for %d-digit ints.\n", digits)
	max := int(math.Pow10(digits)) - 1
	format := fmt.Sprintf("%%0%dd", digits)
	statusInterval := int(math.Pow10(digits - 2))

	memo := map[string]int{} // Memo-ize all findings for efficiency.
	steps := map[int]int{}   // Count how many steps to the answer.
	maxSteps := 0            // Track max number of steps.
	worst := 0               // What is the worst-case iteration?
	canonical := []string{}

	for i := 0; i <= max; i++ {
		s := fmt.Sprintf(format, i)
		n := recurse(memo, s)

		val, _ := steps[n]
		val++ // val is 0 if not found.
		steps[n] = val
		if n > maxSteps {
			maxSteps = n
			worst = i
			assert(maxSteps <= 10)
		}

		if digits > 4 && i%statusInterval == 0 {
			fmt.Printf("%.1e\n", float64(i))
		}

		if n == 0 {
			canonical = append(canonical, s)
		}
	}

	for i := 0; i <= maxSteps; i++ {
		val, _ := steps[i]
		if val == 0 {
			continue
		}
		fmt.Printf("%d ints took %d steps.\n", val, i)
	}
	fmt.Printf("%d was a worst-case recursion which took %d steps.\n", worst, maxSteps)
	fmt.Printf("Canonical: %v\n", canonical)
}

// recurse returns the number of steps from the current int to the canonical
// solution, memo-izing intermediate steps into m as it goes.
func recurse(m map[string]int, s string) int {
	if val, ok := m[s]; ok {
		return val
	}

	n := next(s)
	if n == s { // found the canonical answer
		m[n] = 0
		return 0
	}

	nSteps := recurse(m, n)
	m[s] = nSteps + 1
	return nSteps + 1
}

// next finds the next int in the puzzle sequence based on the current int
func next(s string) string {
	m := map[rune]int{}

	for _, rn := range s {
		val, _ := m[rn]
		m[rn] = val + 1 // val is 0 if not found.
	}

	var b strings.Builder
	for i, rn := range "0123456789" {
		if i == digits-1 {
			break
		}
		val, _ := m[rn] // val is 0 if not found.
		valStr := fmt.Sprintf("%d", val)
		assert(1 == len(valStr))
		b.WriteString(valStr)
	}

	b.WriteString(fmt.Sprintf("%d", len(m))) // Add the digit-count digit.
	assert(digits == b.Len())
	return b.String()
}

// sanity check -- if things are horribly wrong, panic!
func assert(b bool) {
	if !b {
		panic("Assertion failed.")
	}
}
