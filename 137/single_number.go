package _36

//暴力
func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mp := make(map[int]int)
	for _, num := range nums {
		if v, ok := mp[num]; ok {
			if v == 2 {
				delete(mp, num)
			} else {
				mp[num]++
			}

		} else {
			mp[num] = 1
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
	a, b := 0, 0
	for _, v := range nums {
		a = (a ^ v) & ^b
		b = (b ^ v) & ^a
	}
	return a
}
