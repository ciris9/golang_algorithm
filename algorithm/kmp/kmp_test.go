package kmp

import (
	"testing"
)

func TestKmp(t *testing.T) {
	str1and := "aaaabbb"
	str2and := "abb"
	num := Kmp(str1and, str2and)
	t.Logf("%d", num)
}
