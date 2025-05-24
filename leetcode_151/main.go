package main

import (
	"fmt"
	"strings"
)

/*
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.



Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


Constraints:

1 <= s.length <= 104
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.


Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*/

func reverseWords(s string) string {

	s = strings.TrimSpace(s)

	hashMap := make(map[int]int)

	count := 0
	for i := 0; i < len(s); i++ {

		if s[i] == ' ' && s[i+1] != ' ' {
			count++

		} else if count > 0 {

			hashMap[i] = count
			count = 0
		}
	}

	fmt.Println(hashMap)
	//arr := strings.Split(s, " ")
	//
	////slices.Delete(arr, 2, 4)
	//for i, val := range arr {
	//
	//	fmt.Println(i, val)
	//}
	//
	////spaceToRemove := make([]int, 0)
	////
	////start, end := 0, 0
	////for i := 0; i < len(arr); i++ {
	////
	////	if len(arr[i]) == 0 {
	////
	////		start = i
	////	} else {
	////		end = i
	////	}
	////
	////	if start != 0 && end != 0 {
	////
	////	}
	////
	////}
	//
	//slices.Reverse(arr)

	//return strings.Join(arr, " ")

	return ""

}
func main() {

	input := "a good   example"

	fmt.Println(reverseWords(input))

}
