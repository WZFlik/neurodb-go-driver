// Package   net
// @file     msg.go
// @author   Hoss
// @contact  hth146@163.com
// @time     2023/4/23

package neurodb

var ParseOk = 1

type MsgHeadType byte

const (
	MsgTypeParseOk     MsgHeadType = '@'
	MsgTypeSetMsg1     MsgHeadType = '$'
	MsgTypeSetMsg2     MsgHeadType = '#'
	MsgTypeParseObject MsgHeadType = '*'
)

const (
	NEURODB_RETURNDATA byte = 1
	NEURODB_SELECTDB   byte = 2
	NEURODB_EOF        byte = 3
	NEURODB_NODES      byte = 6
	NEURODB_LINKS      byte = 7
	NEURODB_EXIST      byte = 17
	NEURODB_NIL        byte = 18
	NEURODB_RECORD     byte = 19
	NEURODB_RECORDS    byte = 20

	NDB_6BITLEN  byte = 0
	NDB_14BITLEN byte = 1
	NDB_32BITLEN byte = 2
	NDB_ENCVAL   byte = 3

	VO_STRING      byte = 1
	VO_NUM         byte = 2
	VO_STRING_ARRY byte = 3
	VO_NUM_ARRY    byte = 4
	VO_NODE        byte = 5
	VO_LINK        byte = 6
	VO_PATH        byte = 7
	VO_VAR         byte = 8
	VO_VAR_PATTERN byte = 9
)
