// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/24

package dbtype

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
