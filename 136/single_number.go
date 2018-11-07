package _36

//暴力
func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mp := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := mp[num]; ok {
			delete(mp, num)
		} else {
			mp[num] = struct{}{}
		}
	}
	for k := range mp {
		return k
	}
	return 0
}

//位运算
/*
	^ 为异或运算，相同为0，
		a ^ a = 0
		a ^ 0 = a
		a ^ b ^ c = a ^ (b ^ c)
		a^b^a=b
*/
func singleNumber2(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}
