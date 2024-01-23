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

func (c *ColVal) ToInt64() int64 {
	return c.Val.(int64)
}

func (c *ColVal) ToInt32() int32 {
	return c.Val.(int32)
}

func (c *ColVal) ToFloat64() float64 {
	return c.Val.(float64)
}

func (c *ColVal) ToFloat32() float32 {
	return c.Val.(float32)
}

func (c *ColVal) ToInt() int {
	return c.Val.(int)
}

func (c *ColVal) ToString() string {
	return c.Val.(string)
}

func (c *ColVal) ToVal() interface{} {
	return c.Val
}

func (c *ColVal) ToByes() []byte {
	return c.Val.([]byte)
}

func (c *ColVal) ToNode() *Node {
	return c.Val.(*Node)
}

func (c *ColVal) ToLink() *Link {
	return c.Val.(*Link)
}

func (c *ColVal) ToPath() Path {
	return c.Val.([]interface{})
}

func (c *ColVal) ToStrAry() []string {
	return c.Val.([]string)
}

func (c *ColVal) ToNumAry() []float64 {
	return c.Val.([]float64)
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
		return fmt.Sprintf("%v", val)
	}
}
