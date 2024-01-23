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
	Properties  map[string]*ColVal
}

func NewLinks(ID int64, startNodeId int64, endNodeId int64, typ string, properties map[string]*ColVal) *Link {
	return &Link{ID: ID, StartNodeId: startNodeId, EndNodeId: endNodeId, Type: typ, Properties: properties}
}

func (l *Link) GetProperty(key string) (interface{}, bool) {
	val, ok := l.Properties[key]
	if ok {
		return nil, false
	}
	return val.Val, true
}

func (l *Link) getProperty(key string) (*ColVal, bool) {
	val, ok := l.Properties[key]
	if ok {
		return nil, false
	}
	return val, true
}

func (l *Link) GetInt64(key string) (int64, bool) {
	property, b := l.getProperty(key)
	if !b {
		return 0, false
	}
	if _, ok := property.Val.(int64); !ok {
		return 0, false
	}
	return property.ToInt64(), true
}

func (l *Link) GetInt32(key string) (int32, bool) {
	property, b := l.getProperty(key)
	if !b {
		return 0, false
	}
	if _, ok := property.Val.(int32); !ok {
		return 0, false
	}
	return property.ToInt32(), true
}

func (l *Link) GetInt(key string) (int, bool) {
	property, b := l.getProperty(key)
	if !b {
		return 0, false
	}
	if _, ok := property.Val.(int); !ok {
		return 0, false
	}
	return property.ToInt(), true
}

func (l *Link) GetString(key string) (string, bool) {
	property, b := l.getProperty(key)
	if !b {
		return "", false
	}
	if _, ok := property.Val.(string); !ok {
		return "", false
	}
	return property.ToString(), true
}

func (l *Link) GetType() string {
	return l.Type
}

func (l *Link) GetByes(key string) ([]byte, bool) {
	property, b := l.getProperty(key)
	if !b {
		return nil, false
	}
	if _, ok := property.Val.(string); !ok {
		return nil, false
	}
	return property.ToByes(), true
}

func (l *Link) String() string {
	str := fmt.Sprintf("[ID:%d,StartNodeId:%d,EndNodeId:%d,Type:%s:Properties:%v]",
		l.ID, l.StartNodeId, l.EndNodeId, l.Type, l.Properties)
	return str
}
