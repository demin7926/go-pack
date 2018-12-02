package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	var num int
	num = 2
	if isPalindrome(num) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
	num = 121
	if isPalindrome(num) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
	num = 123321
	if isPalindrome(num) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
	num = 123456
	if isPalindrome(num) {
		t.Error("Fail")
	} else {
		t.Log("Pass")
	}
}
