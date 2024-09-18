package main

import "fmt"

func main() {
	// Input slice of integers
	input := []int{1, 2, 3, 3}

	// Output variable to indicate if duplicates are found
	Output := false

	// Map to track seen integers
	temMap := map[int]bool{}

	// Loop through each value in the input slice
	for _, value := range input {
		// Check if the value is already in the map
		if _, has := temMap[value]; has {
			// If the value is already present, set Output to true and break the loop
			Output = true
			break
		} else {
			// If the value is not present, add it to the map
			temMap[value] = true
		}
	}

	// Print the result
	fmt.Println(Output) // Output will be true since there are duplicates
}
