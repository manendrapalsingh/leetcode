package main

import "fmt"

// func main() {

// 	arr := []int{1, 2, 3}

// 	subset := [][]int{}

// 	for i := 0; i < len(arr); i++ {

// 		for j := 0; j < len(arr); j++ {

// 			if i <= len(arr)-j-1 {

// 				subset = append(subset, arr[i:len(arr)-j])
// 			}

// 		}
// 	}

// 	fmt.Println(subset)
// }

// other  way

func main() {
	arr := []int{1, 2, 3}

	subsets := [][]int{}

	for i := 0; i < len(arr); i++ {

		for j := i + 1; j <= len(arr); j++ {

			subset := arr[i:j]
			subsets = append(subsets, subset)

			fmt.Println(subset, i, j)
		}
	}

	fmt.Println("All subsets:", subsets)
}
