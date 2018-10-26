package leetcode

import (
	"fmt"
	"testing"
)

func (l *ListNode) toString() string {
	s := ""
	for t := l; t != nil; t = t.Next {
		s += fmt.Sprintf("%d", t.Val)
	}
	return s
}

func TestSuccessAddTwoNumbers(t *testing.T) {
	// 342 + 465 = 807
	il1, il2, ans1 := [...]int{2, 4, 3}, [...]int{5, 6, 4}, "708"
	l1, l2 := makeListNodeFromSlice(il1[:], il2[:])
	if r1 := AddTwoNumbers(l1, l2); r1.toString() == ans1 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	// 942 + 5168 = 6110
	il3, il4, ans2 := [...]int{2, 4, 9}, [...]int{8, 6, 1, 5}, "0116"
	l3, l4 := makeListNodeFromSlice(il3[:], il4[:])
	if r2 := AddTwoNumbers(l3, l4); r2.toString() == ans2 {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

}
