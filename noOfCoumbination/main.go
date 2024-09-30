package main

import (
	"fmt"
	"math/big"
)

func main() {
	m := 99 // Number of items of type A
	n := 99 // Number of items of type B

	// Create big.Int variables for large number calculations
	a := big.NewInt(int64(m - 1))
	b := big.NewInt(int64(n - 1))

	// Initialize numerator and denominator as big.Int
	numerator := new(big.Int)
	denominator := new(big.Int)
	tempDenominator := new(big.Int)

	// Set numerator to a + b
	numerator.Add(a, b)

	// Determine the number of iterations and initialize the denominator
	var loopShouldBeRun int
	if a.Cmp(b) <= 0 { // Compare a and b
		loopShouldBeRun = m - 1
		tempDenominator.Set(a)
		denominator.Set(a)
	} else {
		loopShouldBeRun = n - 1
		tempDenominator.Set(b)
		denominator.Set(b)
	}

	// Calculate the numerator and denominator iteratively
	for i := 1; i < loopShouldBeRun; i++ {
		fmt.Println(numerator, i, m+n-2-i)

		// Multiply numerator by (m + n - 2 - i)
		numerator.Mul(numerator, big.NewInt(int64(m+n-2-i)))

		// Multiply denominator by (temDenominator - i)
		denominator.Mul(denominator, big.NewInt(int64(tempDenominator.Int64()-int64(i))))
	}

	// Calculate the number of arrangements by dividing the numerator by the denominator
	result := new(big.Int)
	result.Div(numerator, denominator)

	nu := int(result.Int64())

	fmt.Println("Number of arrangements:", result, nu)
}
