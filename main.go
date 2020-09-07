// main recursively solves the puzzle described in the README. By default it
// solves the 8-digit version.
//
// TODO: expand into hexadecimal and explore up to 16-digits!
package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var digits = 8

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
	worst := ""              // What is the worst-case iteration?
	solutions := []string{}

	for i := 0; i <= max; i++ {
		s := fmt.Sprintf(format, i)
		n := recurse(memo, map[string]bool{}, s)

		val, _ := steps[n]
		val++ // val is 0 if not found.
		steps[n] = val
		if n > maxSteps {
			maxSteps = n
			worst = s
			// IDK why we should be able to iterate in fewer steps than the
			// number of digits, but empirically it's always true. Thus when
			// this constraint is violated as a constraint, it's a signal that
			// something is wrong with my implementation.
			if maxSteps > digits {
				fmt.Printf("Warning: %s converged in %d steps.\n", s, n)
			}
		}

		if digits > 5 && i%statusInterval == 0 {
			fmt.Printf("%.1e numbers completed", float64(i))
			if steps[-1] != 0 {
				fmt.Printf(": %d recursions found (so far)", steps[-1])
			}
			fmt.Println(".")
		}

		if n == 0 {
			solutions = append(solutions, s)
			fmt.Println("Found solution!", s)
		}
	}

	for i := 0; i <= maxSteps; i++ {
		val, _ := steps[i]
		if val == 0 {
			continue
		}
		fmt.Printf("%d ints took %d steps.\n", val, i)
	}
	fmt.Printf("%d ints got stuck in loops.\n", steps[-1])
	if maxSteps > 0 {
		fmt.Printf("%s was a worst-case iteration which took %d steps.\n", worst, maxSteps)
	}
	fmt.Printf("Solutions: %v\n", solutions)
}

// recurse returns the number of steps from the current int to a valid solution,
// memo-izing intermediate steps into m as it goes.
func recurse(m map[string]int, seen map[string]bool, s string) int {
	if val, ok := m[s]; ok {
		return val
	}

	n := next(s)
	if n == s { // found a solution
		m[n] = 0
		return 0
	}
	if _, ok := seen[n]; ok {
		return -1 // -1 is a sentinal value for recursion.
	}

	seen[n] = true
	nSteps := recurse(m, seen, n)
	if nSteps == -1 {
		m[s] = -1
		return -1
	}
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
