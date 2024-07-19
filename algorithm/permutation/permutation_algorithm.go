package permutation

import "cmp"

type Permutation[T cmp.Ordered] struct {
	one_ans []T
	isOk    []bool
	ans     [][]T
}

func (p *Permutation[T]) GetPermutation(nums []T) [][]T {
	p.isOk = make([]bool, len(nums))
	p.one_ans = make([]T, 0)
	p.ans = make([][]T, 0)
	p.dfs(0, nums)
	return p.ans
}

func (p *Permutation[T]) dfs(index int, nums []T) {
	if index == len(nums) {
		p.ans = append(p.ans, append([]T(nil), p.one_ans...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if !p.isOk[i] {
			p.one_ans = append(p.one_ans, nums[i])
			p.isOk[i] = true
			p.dfs(index+1, nums)
			p.one_ans = p.one_ans[:len(p.one_ans)-1]
			p.isOk[i] = false
		}
	}
}
