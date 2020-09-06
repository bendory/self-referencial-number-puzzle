# Self-Referencial Number Puzzle

*The PUZZLE and SOLUTION sections below are cut-and-paste copied from the
original puzzle. The empirical analysis is my own.*

## PUZZLE: Self-Referential Number

The first digit of a certain 8-digit integer N is the number of zeroes in the (ordinary, decimal) representation of N.  The second digit is the number of ones; the third, the number of twos; the fourth, the number of threes; the fifth, the number of fours; the sixth, the number of fives; the seventh, the number of sixes; and, finally, the eighth is the total number of distinct digits that appear in N.  What is N?

Source: [Mind-Benders for the
Quarantined](https://momath.org/civicrm/?page=CiviCRM&q=civicrm%2Fevent%2Finfo&id=1620&reset=1)
by Dr. Winkler of the National Museum of Mathmatics.

### SOLUTION

SOLUTION: It's pretty tough to work out N by reasoning.  But there's an easy way!

Take any 8-digit number N0, and transform it according to the puzzle conditions: that is, make a new number N1 by letting the first digit of N1 be the number of 0's in N0, etc.

If by chance N1 = N0, you've solved the puzzle.  But you'd have to be astonishingly lucky for that to happen (the solution to this puzzle is unique).

Not to worry.  Repeat the procedure with N1 to get N2, then N3, until you reach a time T when NT = N(T-1).  Then you're done; NT has the required property.

Let's try this; why not start with 31415926, the first 8 digits of the decimal expansion of pi?

We obtain:

	N0 = 31415926
	N1 = 02111117
	N2 = 15100004
	N3 = 42001104
	N4 = 32102004
	N5 = 31211005
	N6 = 23110105
	N7 = 23110105 = N6 = N (!)

It worked!

Let's try it again with the first 8 digits of e = 2.7182818:

	N0 = 27182818
	N1 = 02200004
	N2 = 50201003
	N3 = 41110105
	N4 = 24001104
	N5 = 32102004
	N6 = 31211005
	N7 = 23110105
	N8 = 23110105 = N7 = N (!)

It worked again!  But why should this work in a reasonable number of steps?  Indeed, why should it work at all?  Couldn't we start to cycle at some point, with NT duplicating some earlier NS with S T-1?

Sadly, I don't know.  This method does not work with all problems of this type.  But it is surprisingly useful — a great weapon to have in your armory.

(Thanks for probabilist Ander Holroyd for creating and contributing this puzzle!)

## Empirical Analysis

Compiling and running `main.go` will check every 8-digit number to empirically
prove that the above iterative approach indeed works for *every* 8-digit number.
But that got me curious -- what about other digit counts?

Empirical results from executing `main.go` follow. For each digit count, the
output includes:

- the number of ints that take *n* steps to iterate to the solution
- the number of ints that got stuck in loops -- which is very common, but
  doens't occur for the 8-digit puzzle
- a worst-case example (a number that took the most number of iterations)
- the canonical solutions to the puzzle for the given number of digits -- which
  turns out to be 0, 1, or 2 solutions depending on the number of digits

Note: there are *10* 1-digit integers because 0 counts.

### 1 Digit

	1 ints took 0 steps.
	9 ints took 1 steps.
	0 ints got stuck in loops.
	0 was a worst-case iteration which took 1 steps.
	Canonical: [1]

### 2 Digits

	100 ints got stuck in loops.
	Canonical: []

### 3 Digits

	1000 ints got stuck in loops.
	Canonical: []

### 4 Digits

	10000 ints got stuck in loops.
	Canonical: []

### 5 Digits

	1 ints took 0 steps.
	359 ints took 1 steps.
	3960 ints took 2 steps.
	10860 ints took 3 steps.
	3456 ints took 4 steps.
	81364 ints got stuck in loops.
	44444 was a worst-case iteration which took 4 steps.
	Canonical: [12104]

### 6 Digits

	2 ints took 0 steps.
	298 ints took 1 steps.
	7500 ints took 2 steps.
	181500 ints took 3 steps.
	497100 ints took 4 steps.
	313600 ints took 5 steps.
	0 ints got stuck in loops.
	000556 was a worst-case iteration which took 5 steps.
	Canonical: [122014 130114]

### 7 Digits

	2 ints took 0 steps.
	348 ints took 1 steps.
	20650 ints took 2 steps.
	483000 ints took 3 steps.
	1493100 ints took 4 steps.
	1403398 ints took 5 steps.
	3064110 ints took 6 steps.
	23184 ints took 7 steps.
	3512208 ints got stuck in loops.
	0000678 was a worst-case iteration which took 7 steps.
	Canonical: [3022003 3103003]

### 8 Digits

	1 ints took 0 steps.
	3359 ints took 1 steps.
	1407840 ints took 2 steps.
	4939200 ints took 3 steps.
	17522400 ints took 4 steps.
	40745460 ints took 5 steps.
	25723446 ints took 6 steps.
	7367026 ints took 7 steps.
	2291268 ints took 8 steps.
	0 ints got stuck in loops.
	00007789 was a worst-case iteration which took 8 steps.
	Canonical: [23110105]

### 9 Digits
