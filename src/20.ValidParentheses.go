// Author			: Cow Cheng
// Created Time		: 07-03-2023
// Last Updated Time: 07-03-2023
// Version			: 1.0

package main

import (
	"fmt"
)

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// 	Open brackets must be closed by the same type of brackets.
// 	Open brackets must be closed in the correct order.
// 	Every close bracket has a corresponding open bracket of the same type.

// Example 1:
//
//	Input: s = "()"
//	Output: true
//
// Example 2:
//
//	Input: s = "()[]{}"
//	Output: true
//
// Example 3:
//
//	Input: s = "(]"
//	Output: false

// Constraints:
//   - 1 <= s.length <= 10^4
//   - s consists of parentheses only '()[]{}'.
func isValid(s string) bool {
	var valArr []string
	var valArrLen int
	var tmpStr string
	var i int

	for i = 0; i < len(s); i++ {
		tmpStr = string(s[i])
		if tmpStr == "(" {
			valArr = append(valArr, ")")
		} else if tmpStr == "[" {
			valArr = append(valArr, "]")
		} else if tmpStr == "{" {
			valArr = append(valArr, "}")
		} else {
			valArrLen = len(valArr)
			if valArrLen == 0 {
				return false
			}
			if tmpStr == valArr[valArrLen-1] {
				valArr = valArr[:valArrLen-1]
			} else {
				return false
			}
		}
	}

	valArrLen = len(valArr)

	if valArrLen == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}
