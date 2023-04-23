// Package   net
// @file     msg.go
// @author   Hoss
// @contact  hth146@163.com
// @time     2023/4/23

package neurodb

// parse the MsgTypeParseObject into an array,HeadIndexxx response the index of array
const  (
	HeadIndexStatus = iota
	HeadIndexCursor // 1
	HeadIndexResults // 2
	HeadIndexAddNodes // 3
	HeadIndexAddLinks // 4
	HeadIndexModifyNodes // 5
	HeadIndexModifyLinks // 6
	HeadIndexDeleteNodes // 7
	HeadIndexDeleteLinks // 8
	HeadIndexBodyLen // 9
	)

var ParseOk = 1

type MsgHeadType byte
const(
	MsgTypeParseOk     MsgHeadType = '@'
	MsgTypeSetMsg1     MsgHeadType = '$'
	MsgTypeSetMsg2     MsgHeadType = '#'
	MsgTypeParseObject MsgHeadType = '*'
)
