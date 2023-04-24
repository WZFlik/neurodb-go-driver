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
		n.ID, n.Labels, n.Properties)
	return str
}
