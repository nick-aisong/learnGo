// 8-31 listing31.go
// 这个示例程序展示如何序列化 JSON 字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 创建一个保存键值对的映射
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// 将这个映射序列化到 JSON 字符串
	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))
}

// G:\GitHub\learnGo\Go语言实战\ch08_标准库\json>go run listing31.go
// {
//  "contact": {
//   "cell": "415.555.5555",
//   "home": "415.333.3333"
//  },
//  "name": "Gopher",
//  "title": "programmer"
// }
