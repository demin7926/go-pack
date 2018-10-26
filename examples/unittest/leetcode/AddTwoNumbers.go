package leetcode

// import (
// 	"fmt"
// )

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var ret, rt *ListNode
	n1, n2, a, s := l1, l2, 0, 0
	for {
		if n1 == nil && n2 == nil {
			if a > 0 {
				rt.Next = &ListNode{Val: 1, Next: nil}
			}
			break
		}
		if n1 != nil && n2 == nil {
			s = n1.Val + a
			n1 = n1.Next
		} else if n1 == nil && n2 != nil {
			s = n2.Val + a
			n2 = n2.Next
		} else {
			s = n1.Val + n2.Val + a
			n1, n2 = n1.Next, n2.Next
		}
		if s > 9 {
			s, a = s-10, 1
		} else {
			a = 0
		}
		if ret == nil {
			rt = &ListNode{Val: s, Next: nil}
			ret = rt
		} else {
			rt.Next = &ListNode{Val: s, Next: nil}
			rt = rt.Next
		}
	}
	return ret
}

func makeListNodeFromSlice(s1, s2 []int) (l1, l2 *ListNode) {
	var t1, t2 *ListNode
	for _, v := range s1 {
		if t1 == nil {
			t1 = &ListNode{Val: v}
			l1 = t1
		} else {
			t1.Next = &ListNode{Val: v}
			t1 = t1.Next
		}
	}
	for _, v := range s2 {
		if t2 == nil {
			t2 = &ListNode{Val: v}
			l2 = t2
		} else {
			t2.Next = &ListNode{Val: v}
			t2 = t2.Next
		}
	}

	return l1, l2
}

// func printNodes(l *ListNode) {
// 	if l == nil {
// 		return
// 	}
// 	for t := l; t != nil; t = t.Next {
// 		fmt.Printf("%d ", t.Val)
// 	}
// 	fmt.Println("")
// }
//
// func main() {
// 	il1, il2 := [...]int{2, 4, 3}, [...]int{5, 6, 4}
// 	l1, l2 := makeListNodeFromSlice(il1[:], il2[:])
// 	printNodes(l1)
// 	printNodes(l2)
// 	r1 := AddTwoNumbers(l1, l2)
// 	printNodes(r1)
// 	fmt.Println("")
//
// 	il3, il4 := [...]int{2, 4, 9}, [...]int{8, 6, 1, 5}
// 	l3, l4 := makeListNodeFromSlice(il3[:], il4[:])
// 	printNodes(l3)
// 	printNodes(l4)
// 	r2 := AddTwoNumbers(l3, l4)
// 	printNodes(r2)
// }
