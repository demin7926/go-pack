package leetcode

import (
	"testing"
)

func isEqual(s1, s2 []int) bool {
	if (s1 == nil) != (s2 == nil) {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}

func TestSuccessTwoSum(t *testing.T) {
	given1, target1 := []int{2, 7, 11, 15}, 9
	ans1 := []int{0, 1}
	if res1 := TwoSum(given1, target1); isEqual(res1, ans1) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	given2, target2 := []int{45, 2, 8, 51, 10}, 61
	ans2 := []int{3, 4}
	if res2 := TwoSum(given2, target2); isEqual(res2, ans2) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	given3, target3 := []int{32, 7, 15, 6}, 22
	ans3 := []int{1, 2}
	if res3 := TwoSum(given3, target3); isEqual(res3, ans3) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

}

func TestFailTwoSum(t *testing.T) {
	given1, target1 := []int{2, 7, 11, 15}, 9
	ans1 := []int{2, 3}
	if res1 := TwoSum(given1, target1); !isEqual(res1, ans1) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	given2, target2 := []int{45, 2, 8, 51, 10}, 61
	ans2 := []int{0, 2}
	if res2 := TwoSum(given2, target2); !isEqual(res2, ans2) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	given3, target3 := []int{32, 7, 15, 6}, 22
	ans3 := []int{0, 3}
	if res3 := TwoSum(given3, target3); !isEqual(res3, ans3) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

}
