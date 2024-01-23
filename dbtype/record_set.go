// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/24

package dbtype

type Record []*ColVal
type Path []interface{}

func (r Record) String() string {
	str := ""
	for i, colVal := range r {
		str += colVal.String()
		if i != len(r)-1 {
			str += ","
		}
	}
	return str
}

func (r Record) ColSize() int {
	return len(r)
}

func (r Record) Col(i int) *ColVal {
	if i > len(r) {
		return nil
	}
	return r[i]
}

type RecordSet struct {
	Labels    []string
	Types     []string
	KeyNames  []string
	Nodes     []*Node
	Links     []*Link
	Records   []Record
	rcdOffset int
}

func (r *RecordSet) GetNodeById(id int64) *Node {
	for _, node := range r.Nodes {
		if node.ID == id {
			return node
		}
	}
	return nil
}

func (r *RecordSet) GetLinkById(id int64) *Link {
	for _, link := range r.Links {
		if link.ID == id {
			return link
		}
	}
	return nil
}

func NewRecordSet() *RecordSet {
	return &RecordSet{rcdOffset: -1}
}

func (r *RecordSet) Next() bool {
	size := len(r.Records)
	if size > 0 && r.rcdOffset+1 < size {
		r.rcdOffset++
		return true
	}
	return false
}

func (r *RecordSet) Record() Record {
	return r.Records[r.rcdOffset]
}

func (r *RecordSet) Err() error {
	// TODO 连接可能会不停迭代数据，此接口用于判断错误
	return nil
}
