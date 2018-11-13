package _260

//暴力
func singleNumber(nums []int) []int {
	mp := make(map[int]struct{})
	for _, one := range nums {
		if _, ok := mp[one]; ok {
			delete(mp, one)
		} else {
			mp[one] = struct{}{}
		}
	}
	s := make([]int, 0)
	for one := range mp {
		s = append(s, one)
	}
	return s
}
