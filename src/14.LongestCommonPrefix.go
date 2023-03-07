// Author			: Cow Cheng
// Created Time		: 07-03-2023
// Last Updated Time: 07-03-2023
// Version			: 1.0

package main

import (
	"fmt"
)

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

// Example 1:
//
//	Input: strs = ["flower","flow","flight"]
//	Output: "fl"
//
// Example 2:
//
//	Input: strs = ["dog","racecar","car"]
//	Output: ""
//	Explanation: There is no common prefix among the input strings.

// Constraints:
//   - 1 <= strs.length <= 200
//   - 0 <= strs[i].length <= 200
//   - strs[i] consists of only lowercase English letters.
func longestCommonPrefix(strs []string) string {
	var strsArrLen int
	var firstStr string
	var i int
	var j int

	strsArrLen = len(strs)

	if strsArrLen == 0 {
		return ""
	}

	firstStr = strs[0]

	if strsArrLen == 1 {
		return firstStr
	}

	for i = 0; i < len(firstStr); i++ {
		for j = 1; j < strsArrLen; j++ {
			if i >= len(strs[j]) || firstStr[i] != strs[j][i] {
				return firstStr[:i]
			}
		}
	}

	return firstStr
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
