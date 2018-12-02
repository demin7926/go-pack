package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	if romanToInt("III") == 3 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	if romanToInt("IV") == 4 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	if romanToInt("IX") == 9 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	if romanToInt("LVIII") == 58 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	if romanToInt("MCMXCIV") == 1994 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}
