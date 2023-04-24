// @file     result_set.go
// @author   Hoss
// @contact  hth146@163.com
// @time     2023/4/23

package dbtype

import (
	"strconv"
)

// parse the MsgTypeParseObject into an array,HeadIndexxx response the index of array
const (
	HeadIndexStatus      = iota
	HeadIndexCursor      // 1
	HeadIndexResults     // 2
	HeadIndexAddNodes    // 3
	HeadIndexAddLinks    // 4
	HeadIndexModifyNodes // 5
	HeadIndexModifyLinks // 6
	HeadIndexDeleteNodes // 7
	HeadIndexDeleteLinks // 8
	HeadIndexBodyLen     // 9
)

type ResultSet interface {
	Next() bool
	Record() (interface{}, error)
}

type ResultInfo struct {
	Status      int
	Cursor      int
	Results     int
	AddNodes    int
	AddLinks    int
	ModifyNodes int
	ModifyLinks int
	DeleteNodes int
	DeleteLinks int
	Msg         string
	BodyLen     int
}

func NewResult() *resultSet {
	return &resultSet{}
}

type resultSet struct {
	ResultInfo
	RecordSet *RecordSet
}

func (r *resultSet) Next() bool {
	panic("implement me")
}

func (r *resultSet) Record() (interface{}, error) {
	panic("implement me")
}

func (r *resultSet) ParseInfo(head []string) error {
	//HeadIndexStatus = 0
	status, err := strconv.Atoi(head[HeadIndexStatus])
	r.Status = status
	if err != nil {
		return err
	}
	//HeadIndexCursor // 1
	cursor, err := strconv.Atoi(head[HeadIndexCursor])
	r.Cursor = cursor
	if err != nil {
		return err
	}
	//HeadIndexResults // 2
	results, err := strconv.Atoi(head[HeadIndexResults])
	r.Results = results
	if err != nil {
		return err
	}

	//HeadIndexAddNodes // 3
	nodes, err := strconv.Atoi(head[HeadIndexAddNodes])
	r.AddNodes = nodes
	if err != nil {
		return err
	}

	//HeadIndexAddLinks // 4
	links, err := strconv.Atoi(head[HeadIndexAddLinks])
	r.AddLinks = links
	if err != nil {
		return err
	}

	//HeadIndexModifyNodes // 5
	modifyNodes, err := strconv.Atoi(head[HeadIndexModifyNodes])
	r.ModifyNodes = modifyNodes
	if err != nil {
		return err
	}

	//HeadIndexModifyLinks // 6
	modifyLinks, err := strconv.Atoi(head[HeadIndexModifyLinks])
	r.ModifyLinks = modifyLinks
	if err != nil {
		return err
	}
	//HeadIndexDeleteNodes // 7
	deleteNodes, err := strconv.Atoi(head[HeadIndexDeleteNodes])
	r.DeleteNodes = deleteNodes
	if err != nil {
		return err
	}

	//HeadIndexDeleteLinks // 8
	deleteLinks, err := strconv.Atoi(head[HeadIndexDeleteLinks])
	r.DeleteLinks = deleteLinks
	if err != nil {
		return err
	}

	//HeadIndexDeleteLinks // 8
	bodyLen, err := strconv.Atoi(head[HeadIndexBodyLen])
	r.BodyLen = bodyLen
	if err != nil {
		return err
	}
	return nil
}
