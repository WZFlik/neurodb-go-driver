# neurodb-go-driver

neurodb golang驱动

# 版本：


​	v0.0.2
# 使用方式
- go get github.com/WZFlik/neurodb-go-driver
- import "github.com/WZFlik/neurodb-go-driver/neurodb"
- 开始编码
```go
   package main

import "github.com/WZFlik/neurodb-go-driver/neurodb"

func main() {
	db, err := neurodb.Open("127.0.0.1", 8839)
	if err != nil {
		panic(err)
	}
	resultSet, err := db.ExecuteQuery("match (n) return n")
	//traversal resultSet
	_ = resultSet
}

```

## 示例代码

```go


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

```



## 运行结果

```go
name: 小米科技有限责任公司
name: 阿里巴巴集团控股有限公司
name: 百度在线网络技术有限公司
name: 雷某
name: 马某
name: 李某
name: 陈某
name: 张某
name: 王某
name: 赵某
[ID:3,Labels:[Person],Properties:map[born:1969 name:雷某]]
[ID:0,StartNodeId:3,EndNodeId:0,Type:WORK:Properties:map[position:CEO]]
[ID:0,Labels:[Company],Properties:map[create_at:2010 name:小米科技有     限责任公司 tagline:致力于让全球每个人都能享受科技带来的美好生活]]
[ID:4,Labels:[Person],Properties:map[born:1964 name:马某]]
[ID:12,StartNodeId:4,EndNodeId:5,Type:FRIEND:Properties:map[weight:10]]
[ID:5,Labels:[Person],Properties:map[born:1968 name:李某]]
[ID:4,Labels:[Person],Properties:map[born:1964 name:马某]]
[ID:2,StartNodeId:4,EndNodeId:1,Type:WORK:Properties:map[position:CEO]]
[ID:1,Labels:[Company],Properties:map[create_at:1999 name:阿里巴巴集     团控股有限公司 tagline:旨在构建未来的商业基础设施]]
[ID:5,Labels:[Person],Properties:map[born:1968 name:李某]]
[ID:4,StartNodeId:5,EndNodeId:2,Type:WORK:Properties:map[position:CEO]]
[ID:2,Labels:[Company],Properties:map[create_at:2000 name:百度在线网     络技术有限公司 tagline:是拥有强大互联网基础的领先AI公司]]
[ID:6,Labels:[Person],Properties:map[born:1960 name:陈某]]
[ID:11,StartNodeId:6,EndNodeId:5,Type:FRIEND:Properties:map[weight:5]]
[ID:5,Labels:[Person],Properties:map[born:1968 name:李某]]
[ID:6,Labels:[Person],Properties:map[born:1960 name:陈某]]
[ID:1,StartNodeId:6,EndNodeId:0,Type:WORK:Properties:map[position:员 工]]
[ID:0,Labels:[Company],Properties:map[create_at:2010 name:小米科技有     限责任公司 tagline:致力于让全球每个人都能享受科技带来的美好生活]]
[ID:7,Labels:[Person],Properties:map[born:1967 name:张某]]
[ID:10,StartNodeId:7,EndNodeId:6,Type:FRIEND:Properties:map[weight:8]]
[ID:6,Labels:[Person],Properties:map[born:1960 name:陈某]]
[ID:7,Labels:[Person],Properties:map[born:1967 name:张某]]
[ID:9,StartNodeId:7,EndNodeId:4,Type:FRIEND:Properties:map[weight:4]]
[ID:4,Labels:[Person],Properties:map[born:1964 name:马某]]
[ID:7,Labels:[Person],Properties:map[born:1967 name:张某]]
[ID:6,StartNodeId:7,EndNodeId:2,Type:WORK:Properties:map[position:员 工]]
[ID:2,Labels:[Company],Properties:map[create_at:2000 name:百度在线网     络技术有限公司 tagline:是拥有强大互联网基础的领先AI公司]]
[ID:8,Labels:[Person],Properties:map[born:1965 name:王某]]
[ID:7,StartNodeId:8,EndNodeId:3,Type:FRIEND:Properties:map[weight:2]]
[ID:3,Labels:[Person],Properties:map[born:1969 name:雷某]]
[ID:8,Labels:[Person],Properties:map[born:1965 name:王某]]
[ID:5,StartNodeId:8,EndNodeId:2,Type:WORK:Properties:map[position:员 工]]
[ID:2,Labels:[Company],Properties:map[create_at:2000 name:百度在线网     络技术有限公司 tagline:是拥有强大互联网基础的领先AI公司]]
[ID:9,Labels:[Person],Properties:map[born:1952 name:赵某]]
[ID:8,StartNodeId:9,EndNodeId:3,Type:FRIEND:Properties:map[weight:1]]
[ID:3,Labels:[Person],Properties:map[born:1969 name:雷某]]
[ID:9,Labels:[Person],Properties:map[born:1952 name:赵某]]
[ID:3,StartNodeId:9,EndNodeId:1,Type:WORK:Properties:map[position:员 工]]
[ID:1,Labels:[Company],Properties:map[create_at:1999 name:阿里巴巴集     团控股有限公司 tagline:旨在构建未来的商业基础设施]]

```


# 更多信息

见[neurodb官网](http://neurodb.org/)

