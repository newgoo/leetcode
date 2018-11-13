package _57

import (
	"testing"

	"github.com/astaxie/beego"
)

func Test_insert(t *testing.T) {
	ls := toList([][]int{{-10, -2}, {11, 12}, {13, 14}})
	in := Interval{1, 7}
	beego.Info(insert(ls, in))
}

func toList(ls [][]int) []Interval {
	v := make([]Interval, len(ls))
	for index, one := range ls {
		if len(one) == 2 {
			v[index].Start = one[0]
			v[index].End = one[1]
		}
	}
	return v
}
