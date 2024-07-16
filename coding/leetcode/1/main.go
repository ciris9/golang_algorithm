package mian

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
