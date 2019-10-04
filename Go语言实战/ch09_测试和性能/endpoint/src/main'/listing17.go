// 9-17 listing17.go
// 这个示例程序实现了简单的网络服务
package main

import (
	"log"
	"net/http"

	"handlers"
)

// main 是应用程序的入口
func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
