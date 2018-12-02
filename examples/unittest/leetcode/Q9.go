package leetcode

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l, m := len(s), len(s)/2
	if m == 0 {
		return true
	}
	for i := 0; i < m; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
