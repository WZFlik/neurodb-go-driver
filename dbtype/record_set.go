// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/24

package dbtype

type RecordSet struct {
	Labels   []string
	Types    []string
	KeyNames []string
	Nodes    []*Node
	Links    []*Link
	Records  [][]*ColVal
}

func (s RecordSet) GetNodeById(id int64) *Node {
	for _, node := range s.Nodes {
		if node.ID == id {
			return node
		}
	}
	return nil
}

func (s RecordSet) GetLinkById(id int64) *Link {
	for _, link := range s.Links {
		if link.ID == id {
			return link
		}
	}
	return nil
}

func NewRecordSet() *RecordSet {
	return &RecordSet{}
}
