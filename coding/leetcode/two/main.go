package two

import (
	"math"
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
