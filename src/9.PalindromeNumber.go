// Author			: Cow Cheng
// Created Time		: 01-03-2023
// Last Updated Time: 01-03-2023
// Version			: 1.0

package main

import (
	"fmt"
)

// Given an integer x, return true if x is a palindrome, and false otherwise.

// Example 1:
//
//	Input: x = 121
//	Output: true
//	Explanation: 121 reads as 121 from left to right and from right to left.
//
// Example 2:
//
//	Input: x = -121
//	Output: false
//	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//
// Example 3:
//
//	Input: x = 10
//	Output: false
//	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Constraints:
//   - -2^31 <= x <= 2^31 - 1

// Follow up: Could you solve it without converting the integer to a string?
func isPalindrome(x int) bool {
	var compList []int
	var compListLen int
	var i int
	var j int

	if x < 0 {
		return false
	}

	for i = x; i != 0; i /= 10 {
		compList = append(compList, i%10)
	}

	compListLen = len(compList) - 1

	for j = 0; j <= compListLen; j++ {
		if compList[j] != compList[compListLen] {
			return false
		}
		compListLen--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
