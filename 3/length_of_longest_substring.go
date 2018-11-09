package _3

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	max := 0
	for i := 0; i < len(s); i++ {
		var j = i
	loop:
		for j = i; j < len(s); j++ {
			if s[i] == s[j] && i != j {
				break
			}
			for m := i; m < j; m++ {
				if s[m] == s[j] {
					break loop
				}
			}
		}
		if v := j - i; max < v {
			max = v
		}
	}
	return max
}

// 超高效率
func lengthOfLongestSubstring2(s string) int {
	max := 0
	dict := make(map[uint8]int, 256)
	//dict := [256]int{}
	//for i := range dict {
	//	dict[i] = 0 // 先设置所有的字符都没有见过
	//}
	left := 0
	for i := 0; i < len(s); i++ {
		if dict[s[i]] >= left {
			left = dict[s[i]]
		}
		if i-left+1 > max {
			max = i - left + 1
		}
		dict[s[i]] = i + 1
		println(string(s[i]), dict[s[i]], i, left)
	}
	return max
}
