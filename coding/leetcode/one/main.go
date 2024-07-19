package one

var one_ans = make([]int, 0)
var isOk = make([]bool, 0)
var ans = make([][]int, 0)

func permute(nums []int) [][]int {
	isOk = make([]bool, len(nums))
	one_ans = make([]int, 0)
	ans = make([][]int, 0)
	dfs(0, nums)
	return ans
}

func dfs(index int, nums []int) {
	if index == len(nums) {
		ans = append(ans, append([]int(nil), one_ans...))
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

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] == 2 && nums[right] == 0 {
			nums[left], nums[right] = nums[right], nums[right]
			left++
			right--
		} else if nums[left] == 2 && nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[right]
			right--
		} else if nums[left] != 2 && nums[right] == 0 {
			nums[left], nums[right] = nums[right], nums[right]
			left++
		} else if nums[left] == 0 {
			left++
		} else if nums[right] == 2 {
			right--
		} else if nums[left] != 0 && nums[right] != 2 {
			now := left
			for nums[now] != 2 && now < right {
				now++
			}
			if now != right {
				nums[now], nums[right] = nums[right], nums[now]
				right--
			}
		}
	}
}

func majorityElement(nums []int) (ans int) {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return
}

func singleNumber(nums []int) (ans int) {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return
}
