package two

import (
	"math"
	"strconv"
	"strings"
)

var ans [][]int
var one_ans []int
var isOk []bool
var isExist map[int]struct{}

func permuteUnique(nums []int) [][]int {
	ans = make([][]int, 0)
	one_ans = make([]int, 0)
	isOk = make([]bool, len(nums))
	isExist = make(map[int]struct{})
	dfs(0, nums)
	return ans
}

func dfs(index int, nums []int) {
	if index == len(nums) {
		cal := caculate(one_ans)
		if _, ok := isExist[cal]; !ok {
			ans = append(ans, append([]int(nil), one_ans...))
			isExist[cal] = struct{}{}
		}
		return
	}
	for i := 0; i < len(nums); i++ {
		if !isOk[i] {
			one_ans = append(one_ans, nums[i])
			isOk[i] = true
			dfs(index+1, nums)
			one_ans = one_ans[:len(one_ans)-1]
			isOk[i] = false
		}
	}
}

func caculate(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res += nums[i] * int(math.Pow10(i))
	}
	return res
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func myAtoi(s string) int {
	trimLefts := strings.TrimLeft(s, " ")
	if trimLefts == "" {
		return 0
	}
	var isBelowZero bool
	if trimLefts[0] == '-' {
		isBelowZero = true
		trimLefts = trimLefts[1:]
	} else if trimLefts[0] == '+' {
		isBelowZero = false
		trimLefts = trimLefts[1:]
	}
	if trimLefts == "" {
		return 0
	}
	trimLefts = strings.TrimLeft(trimLefts, "0")
	for index, trimLeft := range trimLefts {
		if trimLeft < '0' || trimLeft > '9' {
			trimLefts = trimLefts[:index]
			break
		}
	}
	i, _ := strconv.ParseInt(trimLefts, 10, 64)
	if isBelowZero {
		i *= -1
	}
	if i > (1<<31 - 1) {
		return 1<<31 - 1
	} else if i < -(1 << 31) {
		return -(1 << 31)
	}
	return int(i)
}

func jump(nums []int) int {
	step, maxIndex, end := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxIndex = max(maxIndex, i+nums[i])
		if end == i {
			end = maxIndex
			step++
		}
	}
	return step
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

// 假设对于某个数n来说，他的最小数量是m，那么对于n+1来讲就是
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	first, second := head, head.Next
	for first.Next != nil && second.Next != nil && first.Next.Next != nil && second.Next.Next.Next != nil {
		nextFirst := first.Next.Next
		nextSecond := second.Next.Next
		first.Next = nextFirst
		second.Next = first
		first, second = nextSecond, nextFirst
	}
	return head.Next
}
