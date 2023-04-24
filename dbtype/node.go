// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/24

package dbtype

import "fmt"

type Node struct {
	ID         int64
	Labels     []string
	Properties map[string]*ColVal
}

func (n Node) String() string {
	str := fmt.Sprintf("[ID:%d,Labels:%v,Properties:%v]",
		n.ID, toString(n.Labels), toString(n.Properties))
	return str
}

func toString(val interface{}) string {
	switch val.(type) {
	case map[string]*ColVal:
		return prop(val.(map[string]*ColVal)).String()
	case []string:
		return strAry(val.([]string)).String()
	}
	return "#!$"
}
