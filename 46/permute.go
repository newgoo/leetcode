package _46

//8ms
func permute(nums []int) [][]int {
	return p(nil, nums)
}

func p(s1 []int, s2 []int) [][]int {
	if len(s2) == 1 {
		return [][]int{append(s1, s2[0])}
	}
	ls := make([][]int, 0)
	for i := 0; i < len(s2); i++ {
		s := append(s1, s2[i])
		last := append(append([]int{}, s2[:i]...), s2[i+1:]...)
		v := p(s, last)
		ls = append(ls, v...)
	}
	return ls
}
