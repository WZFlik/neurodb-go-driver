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
	RecordSet() *RecordSet
	Next() bool
	Record() Record
	Err() error // 迭代过程中判断Error
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
	firstErr    error
}

func NewResult() *resultSet {
	return &resultSet{}
}

type resultSet struct {
	ResultInfo
	Data     *RecordSet
	firstErr error
}

func (r *resultSet) RecordSet() *RecordSet {
	return r.Data
}

func (r *resultSet) Next() bool {
	return r.Data.Next()
}

// 返回一行记录，按照cypher指令后面的类型排序
// match (n)-[r]->(m) return n,r,m， 则record包含了 Node(n)、Link(r)、Node(m)，这三个组成一条记录
func (r *resultSet) Record() Record {
	return r.Data.Record()
}

func (r *resultSet) Err() error {
	return r.Data.Err()
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
	return nil
}
