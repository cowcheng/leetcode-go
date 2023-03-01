// Author			: Cow Cheng
// Created Time		: 01-03-2023
// Last Updated Time: 01-03-2023
// Version			: 1.0

package main

import (
	"fmt"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
//
//	Input: nums = [2,7,11,15], target = 9
//	Output: [0,1]
//	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//
// Example 2:
//
//	Input: nums = [3,2,4], target = 6
//	Output: [1,2]
//
// Example 3:
//
//	Input: nums = [3,3], target = 6
//	Output: [0,1]

// Constraints:
//   - 2 <= nums.length <= 10^4
//   - -10^9 <= nums[i] <= 10^9
//   - -10^9 <= target <= 10^9
//   - Only one valid answer exists.

// Follow-up: Can you come up with an algorithm that is less than O(n^2) time complexity?
func twoSum(nums []int, target int) []int {
	var numsLen int
	var i int
	var j int

	numsLen = len(nums)

	for i = 0; i < numsLen; i++ {
		for j = i + 1; j < numsLen; j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
