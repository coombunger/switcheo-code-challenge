package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Direct Calculation
Time: O(1), Space: O(1)
*/
func sum_to_n_a(n int) int {
	if n < 0 {
		return 0
	}

	return n * (n + 1) / 2
}

/*
Iterative Solution
Time: O(n), Space: O(1)
*/
func sum_to_n_b(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

/*
Recursive Solution
Time: O(n), Space: O(n)
*/
func sum_to_n_c(n int) int {
	if n <= 0 {
		return 0
	} else {
		return n + sum_to_n_c(n-1)
	}
}

/*
Testing
Expects one integer as cli argument.
*/
func main() {
	n, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("N: %d, A: %d, B: %d, C:%d\n",
		n,
		sum_to_n_a(n),
		sum_to_n_b(n),
		sum_to_n_c(n))
}
