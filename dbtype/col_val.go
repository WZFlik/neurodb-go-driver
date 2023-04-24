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

func (c *ColVal) String() string {
	if c.Val == nil {
		return "nil"
	}
	val := c.Val
	switch val.(type) {
	case *Node:
		return val.(*Node).String()
	case *Link:
		return val.(*Link).String()
	default:
		return fmt.Sprintf("%v",val)
	}
}
