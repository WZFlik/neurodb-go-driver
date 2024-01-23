// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/24

package dbtype

import (
	"fmt"
)

type Node struct {
	ID         int64
	Labels     []string
	Properties map[string]*ColVal
}

func (n *Node) GetProperty(key string) (interface{}, bool) {
	val, ok := n.Properties[key]
	if !ok {
		return nil, false
	}
	return val.Val, true
}

func (n *Node) getProperty(key string) (*ColVal, bool) {
	val, ok := n.Properties[key]
	if !ok {
		return nil, false
	}
	return val, true
}

func (n *Node) GetInt64(key string) (int64, bool) {
	property, b := n.getProperty(key)
	if !b {
		return 0, false
	}
	if _, ok := property.Val.(int64); !ok {
		return 0, false
	}
	return property.ToInt64(), true
}

func (n *Node) GetInt32(key string) (int32, bool) {
	property, b := n.getProperty(key)
	if !b {
		return 0, false
	}
	if _, ok := property.Val.(int32); !ok {
		return 0, false
	}
	return property.ToInt32(), true
}

func (n *Node) GetInt(key string) (int, bool) {
	property, b := n.getProperty(key)
	if !b {
		return 0, false
	}
	if _, ok := property.Val.(int); !ok {
		return 0, false
	}
	return property.ToInt(), true
}

func (n *Node) GetString(key string) (string, bool) {
	property, b := n.getProperty(key)
	if !b {
		return "", false
	}
	if _, ok := property.Val.(string); !ok {
		return "", false
	}
	return property.ToString(), true
}

func (n *Node) GetLabels() []string {
	return n.Labels
}

func (n *Node) GetVal(key string) (interface{}, bool) {
	property, b := n.getProperty(key)
	if !b {
		return nil, false
	}
	return property.ToVal(), true
}

func (n *Node) GetByes(key string) ([]byte, bool) {
	property, b := n.getProperty(key)
	if !b {
		return nil, false
	}
	if _, ok := property.Val.(string); !ok {
		return nil, false
	}
	return property.ToByes(), true
}

func (n *Node) String() string {
	str := fmt.Sprintf("[ID:%d,Labels:%v,Properties:%v]",
		n.ID, n.Labels, n.Properties)
	return str
}
