// Package net
// @author  Hoss
// @contact hth146@163.com
// @time    2023/4/23

package neurodb

import (
	"github.com/WZFlik/neurodb-go-driver/dbtype"
	"sync"
)

type NeuroDBDriver struct {
	conn  NeuroDBConn
	mutex sync.Mutex
}

func Open(addr string, port int) (*NeuroDBDriver, error) {
	conn, err := NewNeuroDBConn(addr, port)
	if err != nil {
		return nil, err
	}
	driver := &NeuroDBDriver{
		conn: conn,
	}
	return driver, nil

}

func (n *NeuroDBDriver) ExecuteQuery(query string) (dbtype.ResultSet, error) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	resultSet, err := n.conn.SendRecv([]byte(query))
	if err != nil {
		return nil, err
	}
	return resultSet, nil
}
