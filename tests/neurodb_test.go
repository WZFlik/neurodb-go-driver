// Package   neurodb
// @file     neurodb_test.go
// @author   Hoss
// @contact  hth146@163.com
// @time     2023/4/23

package tests

import (
	"fmt"
	"github.com/WZFlik/neurodb-go-driver/dbtype"
	"github.com/WZFlik/neurodb-go-driver/neurodb"
	"testing"
)

// If there are any bugs or requirements, please contact us
// go driver github.com/WZFlik/neurodb-go-driver/neurodb
//
//	http://neurodb.org
func TestDriver(t *testing.T) {
	db, err := neurodb.Open("127.0.0.1", 8839)
	if err != nil {
		t.Error(err)
	}

	resultSet, err := db.ExecuteQuery("match (n) return n")
	if err != nil {
		t.Error(err)
	}
	for resultSet.Next() {
		record := resultSet.Record()
		for i := 0; i < record.ColSize(); i++ {
			var col *dbtype.ColVal = record.Col(i)
			var node *dbtype.Node = col.ToNode()
			property, exist := node.GetString("name")
			fmt.Println("name:", property)
			_ = node.GetLabels()
			born, exist := node.GetInt("born")
			if !exist {
				continue
			}
			fmt.Println("born:", born)
		}
	}

	resultSet, err = db.ExecuteQuery("match (n)-[r]->(m) return n,r,m")
	if err != nil {
		t.Error(err)
	}
	for resultSet.Next() {
		record := resultSet.Record()
		// Record contains multiple columns
		for i := 0; i < record.ColSize(); i++ {
			var val *dbtype.ColVal = record.Col(i)
			switch byte(val.Type) {
			case neurodb.VO_NODE:
				var node *dbtype.Node
				node = val.ToNode()
				fmt.Println(node)
			case neurodb.VO_LINK:
				link := val.ToLink()
				fmt.Println(link)
			case neurodb.VO_PATH:
				// []interface{}
				path := val.ToPath()
				fmt.Println(path)
			case neurodb.VO_STRING:
				toString := val.ToString()
				fmt.Println(toString)
			case neurodb.VO_NUM:
				toFloat64 := val.ToFloat64()
				fmt.Println(toFloat64)
			case neurodb.VO_STRING_ARRY:
				strAry := val.ToStrAry()
				_ = strAry
			case neurodb.VO_NUM_ARRY:
				numAry := val.ToNumAry()
				_ = numAry
			default:
				panic("Unknown type Error")
			}
		}
	}
}
