package neurodb

import (
	"errors"
	"fmt"
	"github.com/WZFlik/neurodb-go-driver/dbtype"
	"strconv"
	"unicode"
)

var ErrUnknownType = errors.New("Error Type")

// TODO Hoss consider using bytes directly
// utf8 bytes
func DeserializeRecordSet(s []byte) (*dbtype.RecordSet, error) {
	buffer := newDecMsgBuf(s)
	recordSet := dbtype.NewRecordSet()
	if buffer.DeserializeTypeU1() != NEURODB_RETURNDATA {
		return nil, ErrUnknownType
	}
	/*读取labels、types、keyNames列表*/
	recordSet.Labels = buffer.DeserializeStringList()
	recordSet.Types = buffer.DeserializeStringList()
	recordSet.KeyNames = buffer.DeserializeStringList()
	/*读取节点列表*/
	if buffer.DeserializeTypeU1() != NEURODB_NODES {
		return nil, ErrUnknownType
	}
	nodesCnt := buffer.DeserializeUint()
	recordSet.Nodes = make([]*dbtype.Node, nodesCnt)
	for i := 0; i < nodesCnt; i++ {
		node := buffer.DeserializeCNode(recordSet.Labels, recordSet.KeyNames)
		recordSet.Nodes[i] = node
	}
	/*读取关系列表*/
	if buffer.DeserializeTypeU1() != NEURODB_LINKS {
		return nil, ErrUnknownType
	}
	linksCnt := buffer.DeserializeUint()
	recordSet.Links = make([]*dbtype.Link, linksCnt)
	for i := 0; i < linksCnt; i++ {
		l := buffer.DeserializeClink(recordSet.Types, recordSet.KeyNames)
		recordSet.Links[i] = l
	}
	/*读取return结果集列表*/
	if buffer.DeserializeTypeU1() != NEURODB_RECORDS {
		return nil, ErrUnknownType
	}
	recordsCnt := buffer.DeserializeUint()
	for i := 0; i < recordsCnt; i++ {
		var typ, colCnt int
		if buffer.DeserializeTypeU1() != NEURODB_RECORD {
			return nil, ErrUnknownType
		}
		colCnt = buffer.DeserializeUint()
		rcd := make([]*dbtype.ColVal, 0, colCnt)
		for i := 0; i < colCnt; i++ {
			typ = int(buffer.DeserializeTypeU1())
			val := dbtype.NewColVal()
			val.Type = typ
			switch byte(typ) {
			case NEURODB_NIL:
			// DO NOTHING
			case VO_NODE:
				id := buffer.DeserializeUint()
				node := recordSet.GetNodeById(int64(id))
				//if node == nil {
				//	panic("Node not exist")
				//}
				val.Val = node
			case VO_LINK:
				id := int64(buffer.DeserializeUint())
				l := recordSet.GetLinkById(id)
				//if l == nil {
				//	panic("Link not exist")
				//}
				val.Val = l

			case VO_PATH:
				len := buffer.DeserializeUint()
				path := make([]interface{}, len)
				for i := 0; i < len; i++ {
					id := buffer.DeserializeUint()
					if i%2 == 0 {
						path[i] = recordSet.GetLinkById(int64(id))
						continue
					}
					path[i] = recordSet.GetLinkById(int64(id))
				}
				val.Val = path
			case VO_STRING:
				val.Val = buffer.DeserializeString()
			case VO_NUM:
				fltStr := buffer.DeserializeString()
				floatVal, err := strconv.ParseFloat(fltStr, 64)
				if err != nil {
					panic(err)
				}
				val.Val = floatVal
			case VO_STRING_ARRY:
				aryLen := buffer.DeserializeUint()
				valAry := make([]string, aryLen)
				for i := 0; i < aryLen; i++ {
					valAry[i] = buffer.DeserializeString()
				}
				val.Val = valAry

			case VO_NUM_ARRY:
				aryLen := buffer.DeserializeUint()
				valAry := make([]float64, aryLen)
				for i := 0; i < aryLen; i++ {
					fltStr := buffer.DeserializeString()
					floatVal, err := strconv.ParseFloat(fltStr, 64)
					if err != nil {
						panic(err)
					}
					valAry[i] = floatVal
				}
				val.Val = valAry
			default:
				panic(ErrUnknownType)
			}
			rcd = append(rcd, val)
		}
		recordSet.Records = append(recordSet.Records, rcd)
	}
	if buffer.DeserializeTypeU1() != NEURODB_EOF {
		panic(unicode.Adlam)
	}
	return recordSet, nil
}

type msgBuffer struct {
	Buf    []byte
	offset int
}

// always for receive message
func newMsgBuffer(bufSize int) *msgBuffer {
	return &msgBuffer{
		Buf: make([]byte, bufSize),
	}
}

// always for receive message
func newDecMsgBuf(buf []byte) *msgBuffer {
	return &msgBuffer{
		Buf: buf,
	}

}

func (m *msgBuffer) DeserializeTypeU1() byte {
	m.check()
	tp := m.Buf[m.offset]
	m.offset++
	return tp
}

func (m *msgBuffer) check() {
	if m.offset > len(m.Buf)-1 {
		err := fmt.Errorf("msg buffer out of range offset:%d", m.offset)
		panic(err)
	}
}

func (m *msgBuffer) DeserializeStringList() []string {
	len := m.DeserializeUint()
	labels := make([]string, len)
	for i := 0; i < len; i++ {
		labels[i] = m.DeserializeString()
	}
	return labels
}

func (m *msgBuffer) DeserializeUint() int {
	buf := make([]int, 3)
	bytes := m.readBytes(3)
	for i := 0; i < 3; i++ {
		buf[i] = int(bytes[i])
	}
	return (buf[0]&0x7f)<<14 | (buf[1]&0x7f)<<7 | buf[2]
}

func (m *msgBuffer) DeserializeString() string {
	len := m.DeserializeUint()
	val := string(m.readBytes(len))
	return val

}

func (m *msgBuffer) readBytes(len int) []byte {
	b := make([]byte, len)
	copy(b, m.Buf[m.offset:])
	m.offset += len
	return b
}

func (m *msgBuffer) readByte() byte {
	b := m.Buf[m.offset]
	m.offset++
	return b
}

func (m *msgBuffer) DeserializeCNode(labels []string, keyNames []string) *dbtype.Node {
	var id int64
	var ndsLabels []string
	var kvs map[string]*dbtype.ColVal
	id = int64(m.DeserializeUint())
	ndsLabels = m.DeserializeLabels(labels)
	kvs = m.DeserializeKVList(keyNames)
	return &dbtype.Node{
		ID:         id,
		Labels:     ndsLabels,
		Properties: kvs,
	}
}

func (m *msgBuffer) DeserializeLabels(labels []string) []string {
	listLen := m.DeserializeUint()
	l := make([]string, listLen)
	for i := 0; i < listLen; i++ {
		idx := m.DeserializeUint()
		l[i] = labels[idx]
	}
	return l
}

func (m *msgBuffer) DeserializeKVList(keyNames []string) map[string]*dbtype.ColVal {
	listLen := m.DeserializeUint()
	properties := make(map[string]*dbtype.ColVal)
	for i := 0; i < listLen; i++ {
		idx := m.DeserializeUint()
		key := keyNames[idx]
		typ := m.DeserializeUint()
		colVal := dbtype.NewColVal()
		colVal.Type = typ
		aryLen := 0
		switch byte(typ) {
		case VO_STRING:
			colVal.Val = m.DeserializeString()
		case VO_NUM:
			floatValStr := m.DeserializeString()
			floatVal, err := strconv.ParseFloat(floatValStr, 64)
			if err != nil {
				panic(err)
			}
			colVal.Val = floatVal
		case VO_STRING_ARRY:
			aryLen = m.DeserializeUint()
			vals := make([]string, aryLen)
			aryLen = m.DeserializeUint()
			for j := 0; j < aryLen; j++ {
				vals[j] = m.DeserializeString()
			}
			colVal.Val = vals
		case VO_NUM_ARRY:
			aryLen = m.DeserializeUint()
			vals := make([]float64, aryLen)
			for j := 0; j < aryLen; j++ {
				floatValStr := m.DeserializeString()
				floatVal, err := strconv.ParseFloat(floatValStr, 64)
				if err != nil {
					panic(err)
				}
				vals[j] = floatVal
			}
			colVal.Val = vals
		default:
			panic(ErrUnknownType)
		}
		properties[key] = colVal
	}
	return properties
}

func (m *msgBuffer) DeserializeClink(types []string, keyNames []string) *dbtype.Link {
	var id, hid, tid int64
	var typeIdx int
	var typ string = ""
	var kvs map[string]*dbtype.ColVal
	id = int64(m.DeserializeUint())
	hid = int64(m.DeserializeUint())
	tid = int64(m.DeserializeUint())
	var ty int
	ty = int(m.DeserializeTypeU1())
	if ty == int(NEURODB_EXIST) {
		typeIdx = m.DeserializeUint()
		typ = types[typeIdx]
	} else if ty == int(NEURODB_NIL) {
		// TODO HOUTH empty？
	}
	kvs = m.DeserializeKVList(keyNames)
	l := dbtype.NewLinks(id, hid, tid, typ, kvs)
	return l
}
