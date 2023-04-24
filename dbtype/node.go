// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/24

package dbtype

type Node struct {
	ID         int64
	Labels     []string
	Properties map[string]*ColVal
}

func (n Node) String() string {
	panic("implement me")
}
