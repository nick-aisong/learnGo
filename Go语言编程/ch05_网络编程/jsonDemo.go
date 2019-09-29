// jsonDemo
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Book struct {
		Title       string
		Authors     []string
		Publisher   string
		IsPublished bool
		Price       float32
	}

	gobook := Book{
		Title:       "Go语言编程",
		Authors:     []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		IsPublished: true,
		Price:       9.99,
	}

	b, err := json.Marshal(gobook)
	if err != nil {
		//...
	}

	fmt.Printf("%s\n", b)

	b, err = json.MarshalIndent(gobook, "", "    ")
	if err != nil {
		//...
	}

	fmt.Printf("%s\n", b)
}

//G:\GitHub\learnGo\Go语言编程\ch05_网络编程>go run jsonDemo.go
//{"Title":"Go语言编程","Authors":["XuShiwei","HughLv","Pandaman","GuaguaSong","HanTuo","BertYuan","XuDaoli"],"Publisher":"","IsPublished":true,"Price":9.99}
//{
//    "Title": "Go语言编程",
//    "Authors": [
//        "XuShiwei",
//        "HughLv",
//        "Pandaman",
//        "GuaguaSong",
//        "HanTuo",
//        "BertYuan",
//        "XuDaoli"
//    ],
//    "Publisher": "",
//    "IsPublished": true,
//    "Price": 9.99
//}
