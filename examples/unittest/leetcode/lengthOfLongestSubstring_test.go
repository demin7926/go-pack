package leetcode

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var str string
	var ans, res int

	str, ans = "abcabcbb", 3
	if res = LengthOfLongestSubstring(str); res == ans {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	str, ans = "bbbbb", 1
	if res = LengthOfLongestSubstring(str); res == ans {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	str, ans = "pwwkew", 3
	if res = LengthOfLongestSubstring(str); res == ans {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}
