// Package   errno
// @file     err.go
// @author   Hoss
// @contact  hth146@163.com
// @time     2023/4/23

package errno

import "fmt"

var (
	//ParseOk = 0
	NoMemErr  = NewNeuroError(2, "No_MEM_ERR")
	SyntaxErr = NewNeuroError(3, "SYNTAX_ERR")
	// todo Hoss more error 3 - 29

	// iterator ResultSet Error when traversal data
	ErrResultSetEmpty = NewNeuroError(999, "Has no next record")
	ErrPropNotExist   = NewNeuroError(998, "Property Message does not exists")
)

type NeuroErrr struct {
	code int
	desc string
}

func (n NeuroErrr) Error() string {
	return fmt.Sprintf("[%d]%s", n.code, n.desc)
}

func NewNeuroError(code int, desc string) NeuroErrr {
	return NeuroErrr{
		code: code,
		desc: desc,
	}
}
