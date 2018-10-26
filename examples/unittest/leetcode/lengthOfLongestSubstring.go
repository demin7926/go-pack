package leetcode

import (
	//"flag"
	"fmt"
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	ts, rs := "", ""
	for _, c := range s {
		cs := fmt.Sprintf("%c", c)
		if j := strings.Index(ts, cs); j > -1 {
			if j == len(ts)-1 {
				ts = cs
			} else {
				ts = ts[j+1:] + cs
			}
		} else {
			ts = fmt.Sprintf("%s%c", ts, c)
			if len(rs) < len(ts) {
				rs = ts
			}
		}
	}
	return len(rs)
}
