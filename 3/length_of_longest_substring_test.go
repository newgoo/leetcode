package _3

import (
	"testing"
)

func TestLength_of_longest_substring(t *testing.T) {
	mp := map[string]int{
		//"au":     2,
		//"a":      1,
		//"aaa":    1,
		//"aba":    2,
		//"aab":    2,
		//" ":      1,
		"pwwkew": 3,
	}
	for key, one := range mp {
		r := lengthOfLongestSubstring2(key)
		if one != r {
			t.Error(key, one, r)
		}
		//println(key, one, r)
	}

}
