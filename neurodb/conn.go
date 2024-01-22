// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/23

package neurodb

import (
	"github.com/WZFlik/neurodb-go-driver/dbtype"
	"net"
	"strconv"
	"strings"
	"sync"
)

const (
	DefaultMsgBuffSize = 1024
)

// A connection of neurodb, it's thread safe
type NeuroDBConn interface {
	SendRecv(msg []byte) (dbtype.ResultSet, error)
	Close() error
}

type neuroDBConn struct {
	tcp *net.TCPConn
	*msgBuffer
	mutex sync.Mutex
}

func NewNeuroDBConn(addr string, port int) (NeuroDBConn, error) {
	conn := &neuroDBConn{}
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr+":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	conn.tcp, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	err = conn.tcp.SetKeepAlive(true)
	if err != nil {
		return nil, err
	}
	err = conn.tcp.SetNoDelay(true)
	if err != nil {
		return nil, err
	}
	conn.msgBuffer = newMsgBuffer(DefaultMsgBuffSize)
	return conn, nil
}

func (n *neuroDBConn) send(msg []byte) error {
	_, err := n.tcp.Write(msg)
	if err != nil {
		return err
	}
	return nil
}

func (n *neuroDBConn) recv() (dbtype.ResultSet, error) {
	tp, err := n.readByte()
	if err != nil {
		return nil, err
	}
	resultSet := dbtype.NewResult()

	switch MsgHeadType(tp) {
	case MsgTypeParseOk: // @
		resultSet.Status = ParseOk
	case MsgTypeSetMsg1, MsgTypeSetMsg2: // $
		line, err := n.readLine()
		if err != nil {
			return nil, err
		}
		resultSet.Msg = line
	case MsgTypeParseObject:
		line, err := n.readLine()
		if err != nil {
			return nil, err
		}
		head := strings.Split(line, ",")
		err = resultSet.ParseInfo(head)
		if err != nil {
			return nil, err
		}
		//HeadIndexDeleteLinks // 8
		bodyLen, err := strconv.Atoi(head[dbtype.HeadIndexBodyLen])
		if err != nil {
			return nil, err
		}

		body, err := n.readBytes(bodyLen)
		if err != nil {
			return nil, err
		}
		// TODO Hoss  current line not used is bug?
		line, err = n.readLine()
		if err != nil {
			return nil, err
		}
		_ = line
		resultSet.Data, err = DeserializeRecordSet(body)
		if err != nil {
			return nil, err
		}
	}

	return resultSet, nil
}

func (n *neuroDBConn) readLine() (string, error) {
	str := make([]byte, 0)
	charBuf := make([]byte, 1)
	for {
		_, err := n.tcp.Read(charBuf)
		if err != nil {
			return "", err
		}
		str = append(str, charBuf[0])
		if charBuf[0] == '\n' {
			break
		}
	}

	return strings.ReplaceAll(string(str), "\r\n", ""), nil

}

func (n *neuroDBConn) readBytes(len int) ([]byte, error) {
	buf := make([]byte, len)
	_, err := n.tcp.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (n *neuroDBConn) readByte() (byte, error) {
	msgTp := make([]byte, 1)
	_, err := n.tcp.Read(msgTp)
	if err != nil {
		return 0, err
	}
	return msgTp[0], nil

}

// return a ResultSet or an error
func (n *neuroDBConn) SendRecv(msg []byte) (dbtype.ResultSet, error) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	err := n.send(msg)
	if err != nil {
		return nil, err
	}
	resultSet, err := n.recv()
	if err != nil {
		return nil, err
	}
	return resultSet, nil
}

func (n *neuroDBConn) Close() error {
	return n.tcp.Close()
}
