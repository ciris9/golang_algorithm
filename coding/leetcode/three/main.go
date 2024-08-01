package three

import (
	"container/list"
	"strconv"
)

func isPalindrome(x int) bool {
	formatInt := strconv.FormatInt(int64(x), 10)
	for i := 0; i < len(formatInt)/2; i++ {
		if formatInt[i] != formatInt[len(formatInt)-i-1] {
			return false
		}
	}
	return true
}

/*
'('，')'，'{'，'}'，'['，']'
*/
func isValid(s string) bool {
	stack := list.New()
	for _, ch := range s {
		var top int32
		if stack.Len() != 0 {
			top = stack.Back().Value.(int32)
		}
		switch ch {
		case ')':
			if top == '(' {
				stack.Remove(stack.Back())
			} else {
				return false
			}
		case '}':
			if top == '{' {
				stack.Remove(stack.Back())
			} else {
				return false
			}
		case ']':
			if top == '[' {
				stack.Remove(stack.Back())
			} else {
				return false
			}
		default:
			stack.PushBack(ch)
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}

func rob(nums []int) int {
	if len(nums) <= 0 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(dp)-1]
}

func coindfs(coins []int, amount int) {

}
func coinChange(coins []int, amount int) int {

}
