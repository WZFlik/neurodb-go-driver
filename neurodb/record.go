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

type Node struct {

}

func (n Node) String() string {
	panic("implement me")
}

type Relationship struct {

}

func (r Relationship) String() string {
	panic("implement me")
}

