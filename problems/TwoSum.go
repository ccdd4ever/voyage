//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution.
//see @https://leetcode.com/problems/two-sum/

package problems

import "fmt"

func twoSum(nums []int, target int) []int {
	for index := 0; index < len(nums); index++ {
		temp := target - nums[index]
		for i := index + 1; i < len(nums); i++ {
			if temp == nums[i] {
				return []int{index, i}
			}
		}
	}
	return nil
}

//just for
func main() {
	// fmt.Println(add(-20, 3))
	fmt.Println(twoSum([]int{20, 1, 2}, 3))
}
