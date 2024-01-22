// Package   neurodb
// @file     neurodb_test.go
// @author   Hoss
// @contact  hth146@163.com
// @time     2023/4/23

package tests

import (
	"fmt"
	"github.com/WZFlik/neurodb-go-driver/neurodb"
	"testing"
)

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
		if err := resultSet.Err(); err != nil {
			panic(err)
		}
		record := resultSet.Record()
		fmt.Println(record)
	}
	resultSet, err = db.ExecuteQuery("match (n)-[r]->(m) return n,r,m")

	if err != nil {
		t.Error(err)
	}
	for resultSet.Next() {
		if err := resultSet.Err(); err != nil {
			panic(err)
		}
		record := resultSet.Record()
		fmt.Println(record)
	}
}
