// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/24

package dbtype

import "fmt"

type Link struct {
	ID          int64
	StartNodeId int64
	EndNodeId   int64
	Type        string
	properties  map[string]*ColVal
}

func NewLinks(ID int64, startNodeId int64, endNodeId int64, typ string, properties map[string]*ColVal) *Link {
	return &Link{ID: ID, StartNodeId: startNodeId, EndNodeId: endNodeId, Type: typ, properties: properties}
}

func (l Link) String() string {
	str := fmt.Sprintf("[ID:%d,StartNodeId:%d,EndNodeId:%d,Type:%s:properties:%v]",
		l.ID, l.StartNodeId, l.EndNodeId, l.Type, l.properties)
	return str
}
