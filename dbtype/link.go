// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/24

package dbtype

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

func (r Link) String() string {
	panic("implement me")
}
