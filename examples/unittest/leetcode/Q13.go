package leetcode

import "fmt"

func romanToInt(s string) int {
	m1 := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	m2 := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	ret, l := 0, len(s)
	for i := 0; i < l; i++ {
		c1 := s[i]
		n := m1[c1]
		if i < l-1 {
			c2 := s[i+1]
			sk := fmt.Sprintf("%c%c", c1, c2)
			t1, found := m2[sk]
			if found {
				n = t1
				i = i + 1
			}
		}
		ret += n
	}

	return ret
}
