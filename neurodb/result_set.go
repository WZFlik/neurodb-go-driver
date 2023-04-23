// Package   neurodb
// @file     result_set.go
// @author   Hoss
// @contact  hth146@163.com
// @time     2023/4/23

package neurodb

import "strconv"

type ResultSet interface {
	Next() bool
	Record() (interface{},error)
}


type ResultInfo struct {
	status int
	cursor int
	results int
	addNodes int
	addLinks int
	modifyNodes int
	modifyLinks int
	deleteNodes int
	deleteLinks int
	msg string
	bodyLen int

}

func NewResult() *resultSet {
	return &resultSet{}
}
type resultSet struct {
	ResultInfo
	recordSet []Record
}

func (r resultSet) Next() bool {
	panic("implement me")
}

func (r resultSet) Record() (interface{}, error) {
	panic("implement me")
}

func (r resultSet) parseInfo(head []string) error{
	//HeadIndexStatus = 0
	status, err := strconv.Atoi(head[HeadIndexStatus])
	r.status = status
	if err != nil {
		return err
	}
	//HeadIndexCursor // 1
	cursor, err := strconv.Atoi(head[HeadIndexCursor])
	r.status = cursor
	if err != nil {
		return err
	}
	//HeadIndexResults // 2
	results, err := strconv.Atoi(head[HeadIndexResults])
	r.results = results
	if err != nil {
		return err
	}

	//HeadIndexAddNodes // 3
	nodes, err := strconv.Atoi(head[HeadIndexAddNodes])
	r.addNodes = nodes
	if err != nil {
		return err
	}

	//HeadIndexAddLinks // 4
	links, err := strconv.Atoi(head[HeadIndexAddLinks])
	r.addLinks = links
	if err != nil {
		return err
	}

	//HeadIndexModifyNodes // 5
	modifyNodes, err := strconv.Atoi(head[HeadIndexModifyNodes])
	r.modifyNodes = modifyNodes
	if err != nil {
		return err
	}

	//HeadIndexModifyLinks // 6
	modifyLinks, err := strconv.Atoi(head[HeadIndexModifyLinks])
	r.modifyLinks = modifyLinks
	if err != nil {
		return err
	}
	//HeadIndexDeleteNodes // 7
	deleteNodes, err := strconv.Atoi(head[HeadIndexDeleteNodes])
	r.deleteNodes = deleteNodes
	if err != nil {
		return err
	}

	//HeadIndexDeleteLinks // 8
	deleteLinks, err := strconv.Atoi(head[HeadIndexDeleteLinks])
	r.deleteLinks = deleteLinks
	if err != nil {
		return err
	}

	//HeadIndexDeleteLinks // 8
	bodyLen, err := strconv.Atoi(head[HeadIndexBodyLen])
	r.bodyLen = bodyLen
	if err != nil {
		return err
	}
	return nil
}
