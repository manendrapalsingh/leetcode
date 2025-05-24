package main

import (
	"fmt"
	"strconv"
)

/*
You are given a string num, representing a large integer. Return the largest-valued odd integer (as a string) that is a non-empty substring of num, or an empty string "" if no odd integer exists.

A substring is a contiguous sequence of characters within a string.



Example 1:

Input: num = "52"
Output: "5"
Explanation: The only non-empty substrings are "5", "2", and "52". "5" is the only odd number.
Example 2:

Input: num = "4206"
Output: ""
Explanation: There are no odd numbers in "4206".
Example 3:

Input: num = "35427"
Output: "35427"
Explanation: "35427" is already an odd number.


Constraints:

1 <= num.length <= 105
num only consists of digits and does not contain any leading zeros.
*/

func largestOddNumber(num string) string {

	numInt, _ := strconv.Atoi(num)

	output := ""

	if numInt%2 == 0 {

		oddNumber := 0

		for _, val := range num {

			number, _ := strconv.Atoi(string(val))

			if number > oddNumber {
				oddNumber = number
			}
		}

		if oddNumber != 0 {
			output = strconv.Itoa(oddNumber)
		}

	} else {
		return num
	}

	return output
}
func main() {

	input1 := "52"
	//input2 := "4206"
	//input3 := "35427"

	fmt.Println(largestOddNumber(input1))
	//fmt.Println(largestOddNumber(input2))
	//fmt.Println(largestOddNumber(input3))

}
