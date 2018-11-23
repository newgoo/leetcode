package _46

import (
	"testing"
)

type TS struct {
	input  []int
	result [][]int
}

func Test_Permute(t *testing.T) {
	ts := []TS{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{1, 2, 3, 4}, [][]int{
			{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2},
			{2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1},
			{3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1},
			{4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1},
		}},
	}
	//result := permute([]int{1, 2, 3, 4})
	//s := ""
	//for _, one := range result {
	//	s += "{"
	//	for _, i := range one {
	//		s += fmt.Sprintf("%d,", i)
	//	}
	//	s = s[:len(s)-1]
	//	s += "},"
	//}
	for _, one := range ts {
		result := permute(one.input)
		for i, s := range result {
			for inx, sd := range s {
				if one.result[i][inx] != sd {
					t.Error("不正确")
				}
			}
		}
	}
}
