# neurodb-go-driver

neurodb golang驱动

# 版本：

​	v0.0.1
# 修复问题
1.修正driver module name 修复获取 driver源码时出现module定义路径不匹配的问题

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
```



## 运行结果

```go
[ID:0,Labels:[Company],Properties:map[create_at:2010 name:小米科技有限责任公司 tagline:致力于让全球每个人都能享受科技带来的美好生活]]
[ID:1,Labels:[Company],Properties:map[create_at:1999 name:阿里巴巴集团控股有限公司 tagline:旨在构建未来的商业基础设施]]
[ID:2,Labels:[Company],Properties:map[create_at:2000 name:百度在线网络技术有限公司 tagline:是拥有强大互联网基础的领先AI公司]]
[ID:3,Labels:[Person],Properties:map[born:1969 name:雷某]]
[ID:4,Labels:[Person],Properties:map[born:1964 name:马某]]
[ID:5,Labels:[Person],Properties:map[born:1968 name:李某]]
[ID:6,Labels:[Person],Properties:map[born:1960 name:陈某]]
[ID:7,Labels:[Person],Properties:map[born:1967 name:张某]]
[ID:8,Labels:[Person],Properties:map[born:1965 name:王某]]
[ID:9,Labels:[Person],Properties:map[born:1952 name:赵某]]
[ID:3,Labels:[Person],Properties:map[born:1969 name:雷某]],[ID:0,StartNodeId:3,EndNodeId:0,Type:WORK:properties:map[position:CEO]],[ID:0,Labels:[Company],Properties:map[create_at:2010 name:小米科技有限责任公司 tagline:致力于让全球每个人都能享受科技带来的美好生活]]
[ID:4,Labels:[Person],Properties:map[born:1964 name:马某]],[ID:12,StartNodeId:4,EndNodeId:5,Type:FRIEND:properties:map[weight:10]],[ID:5,Labels:[Person],Properties:map[born:1968 name:李某]]
[ID:4,Labels:[Person],Properties:map[born:1964 name:马某]],[ID:2,StartNodeId:4,EndNodeId:1,Type:WORK:properties:map[position:CEO]],[ID:1,Labels:[Company],Properties:map[create_at:1999 name:阿里巴巴集团控股有限公司 tagline:旨在构建未来的商业基础设施]]
[ID:5,Labels:[Person],Properties:map[born:1968 name:李某]],[ID:4,StartNodeId:5,EndNodeId:2,Type:WORK:properties:map[position:CEO]],[ID:2,Labels:[Company],Properties:map[create_at:2000 name:百度在线网络技术有限公司 tagline:是拥有强大互联网基础的领先AI公司]]
[ID:6,Labels:[Person],Properties:map[born:1960 name:陈某]],[ID:11,StartNodeId:6,EndNodeId:5,Type:FRIEND:properties:map[weight:5]],[ID:5,Labels:[Person],Properties:map[born:1968 name:李某]]
[ID:6,Labels:[Person],Properties:map[born:1960 name:陈某]],[ID:1,StartNodeId:6,EndNodeId:0,Type:WORK:properties:map[position:员工]],[ID:0,Labels:[Company],Properties:map[create_at:2010 name:小米科技有限责任公司 tagline:致力于让全球每个人都能享受科技带来的美好生活]]
[ID:7,Labels:[Person],Properties:map[born:1967 name:张某]],[ID:10,StartNodeId:7,EndNodeId:6,Type:FRIEND:properties:map[weight:8]],[ID:6,Labels:[Person],Properties:map[born:1960 name:陈某]]
[ID:7,Labels:[Person],Properties:map[born:1967 name:张某]],[ID:9,StartNodeId:7,EndNodeId:4,Type:FRIEND:properties:map[weight:4]],[ID:4,Labels:[Person],Properties:map[born:1964 name:马某]]
[ID:7,Labels:[Person],Properties:map[born:1967 name:张某]],[ID:6,StartNodeId:7,EndNodeId:2,Type:WORK:properties:map[position:员工]],[ID:2,Labels:[Company],Properties:map[create_at:2000 name:百度在线网络技术有限公司 tagline:是拥有强大互联网基础的领先AI公司]]
[ID:8,Labels:[Person],Properties:map[born:1965 name:王某]],[ID:7,StartNodeId:8,EndNodeId:3,Type:FRIEND:properties:map[weight:2]],[ID:3,Labels:[Person],Properties:map[born:1969 name:雷某]]
[ID:8,Labels:[Person],Properties:map[born:1965 name:王某]],[ID:5,StartNodeId:8,EndNodeId:2,Type:WORK:properties:map[position:员工]],[ID:2,Labels:[Company],Properties:map[create_at:2000 name:百度在线网络技术有限公司 tagline:是拥有强大互联网基础的领先AI公司]]
[ID:9,Labels:[Person],Properties:map[born:1952 name:赵某]],[ID:8,StartNodeId:9,EndNodeId:3,Type:FRIEND:properties:map[weight:1]],[ID:3,Labels:[Person],Properties:map[born:1969 name:雷某]]
[ID:9,Labels:[Person],Properties:map[born:1952 name:赵某]],[ID:3,StartNodeId:9,EndNodeId:1,Type:WORK:properties:map[position:员工]],[ID:1,Labels:[Company],Properties:map[create_at:1999 name:阿里巴巴集团控股有限公司 tagline:旨在构建未来的商业基础设施]]

```

# 使用

使用 go get命令或者下载源码使用

# 更多信息

[neurodb官网](http://neurodb.org/)

