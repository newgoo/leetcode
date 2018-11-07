package _36

import (
	"testing"

	"github.com/astaxie/beego"
)

func TestSingleNumber(t *testing.T) {
	tm := []int{4, 4, 4, 1}
	beego.Info(singleNumber(tm))
}

func TestSingleNumber2(t *testing.T) {
	tm := []int{4, 1, 1, 1, 2, 2, 2}
	beego.Info(singleNumber2(tm))
}
