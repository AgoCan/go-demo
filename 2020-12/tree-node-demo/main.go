package main

import (
	"encoding/json"
	"fmt"
)

// TreeNode 123
type TreeNode struct {
	ID   int         `json:"id"`
	PID  int         `json:"pid"`
	Name string      `json:"name"`
	List []*TreeNode `json:"list,omitempty"` // omitempty 表示如果空的话,则忽略.
}

func main() {
	test := `[{"id":1,"pid":0,"name":"a"},{"id":2,"pid":1,"name":"b"},{"id":3,"pid":1,"name":"c"},{"id":3,"pid":2,"name":"d"}]`
	var list []*TreeNode // 假设已经赋值进去，从 数据库取出的数据
	if err := json.Unmarshal([]byte(test), &list); err != nil {
		return
	}
	// a.构建map结构的数据.
	data := buildData(list)
	// b.构建树结构
	result := buildTree(0, data)
	temp, _ := json.Marshal(result)
	fmt.Println(string(temp))
}

// 构建parent_id为key的map结构.
func buildData(list []*TreeNode) map[int]map[int]*TreeNode {
	data := make(map[int]map[int]*TreeNode)
	for _, v := range list {
		id := v.ID                   // 主ID
		pid := v.PID                 // 父ID
		if _, ok := data[pid]; !ok { // 如果不存在则创建一个新节点
			data[pid] = make(map[int]*TreeNode)
		}
		data[pid][id] = v
	}
	return data
}

// 构建树的结构.
// a. 判断parent_id是否存在.
// b. 如果parent_id存在继续递归.至到data没有找到parent_id节点的数据.
func buildTree(parentID int, data map[int]map[int]*TreeNode) []*TreeNode {
	node := make([]*TreeNode, 0)
	for id, item := range data[parentID] {
		if data[id] != nil {
			item.List = buildTree(id, data)
		}
		node = append(node, item)
	}
	return node
}
