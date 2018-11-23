package _47

import (
	"testing"

	"github.com/astaxie/beego"
)

func Test_PermuteUnique(t *testing.T) {
	beego.Info(permuteUnique([]int{1, 2, 1, 2}))
}
