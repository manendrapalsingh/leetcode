package main

import "fmt"

// Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

// Example 1:

// Input: x = 2.00000, n = 10
// Output: 1024.00000
// Example 2:

// Input: x = 2.10000, n = 3
// Output: 9.26100
// Example 3:

// Input: x = 2.00000, n = -2
// Output: 0.25000
// Explanation: 2-2 = 1/22 = 1/4 = 0.25

// Constraints:

// -100.0 < x < 100.0
// -231 <= n <= 231-1
// n is an integer.
// Either x is not zero or n > 0.
// -104 <= xn <= 104

func main() {

	x := -1.00000
	n := 2147483648

	var Output float64

	if n > 0 && (x != 1 && x != -1) {
		Output = x
		fmt.Println(1, "===================")

		for i := 1; i < n; i++ {

			Output = Output * x

		}
	} else if n < 0 && (x != 1 && x != -1) {
		fmt.Println(2, "===================")

		Output = 1 / x
		n = -n

		for i := 1; i < n && Output > 0; i++ {

			Output = Output * 1 / x

		}
	} else if n == 0 {
		fmt.Println(3, "===================")

		Output = 1
	} else if x == 1 || x == -1 {
		fmt.Println(4, "===================")

		Output = x

	}

	fmt.Println(Output)

}
