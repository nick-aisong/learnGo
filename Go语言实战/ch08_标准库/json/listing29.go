// 8-29 listing29.go
// 这个示例程序展示如何解码 JSON 字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON 包含要反序列化的样例字符串
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	// 将 JSON 字符串反序列化到 map 变量
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}

// G:\GitHub\learnGo\Go语言实战\ch08_标准库\json>go run listing29.go
// Name: Gopher
// Title: programmer
// Contact
// H: 415.333.3333
// C: 415.555.5555
