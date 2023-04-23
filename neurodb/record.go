// Package   neurodb
// @file     record.go
// @author   Hoss
// @contact  hth146@163.com
// @time     2023/4/23

package neurodb

// TODO Hoss consider using bytes directly
func decodeRecordSet(s string) []Record {
	// TODO Hoss decode record from s
	panic("implement me")
	return nil
}

type Record interface {
	String() string
}

type Vertex struct {

}

func (v Vertex) String() string {
	panic("implement me")
}

type Edge struct {

}

func (e Edge) String() string {
	panic("implement me")
}

