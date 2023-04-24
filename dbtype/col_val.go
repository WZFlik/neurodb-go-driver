// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/24

package dbtype

import (
	"fmt"
)

type ColVal struct {
	Val    interface{}
	Type   int
	AryLen int
}

func (c ColVal) GetNum() int64 {
	return c.Val.(int64)
}

func NewColVal() *ColVal {
	return &ColVal{}
}

func (c ColVal) String() string {
	str := fmt.Sprintf("val:%v,Type:%d", toString(c.Val), c.Type)
	return str
}
