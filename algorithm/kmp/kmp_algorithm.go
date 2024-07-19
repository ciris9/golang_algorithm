package kmp

func initNext(s string) []int {
	if len(s) == 1 {
		return []int{-1}
	}
	next := make([]int, len(s))
	i, j, L := 1, 0, len(s)
	for i < L {
		if s[i] == s[j] {
			j++
			next[i] = j
			i++
		} else {
			if j > 0 {
				j = next[j-1]
			} else {
				next[i] = j
				i++
			}
		}
	}
	return next
}

func Kmp(s, substr string) int {
	next := initNext(substr)
	L := len(s)
	targetL := len(s)
	j := 0
	for i := 0; i < L; i++ {
		if s[i] == substr[j] {
			if j == targetL-1 {
				return i - targetL + 1
			}
			j++
			i++
		} else {
			if j > 0 {
				j = next[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}
