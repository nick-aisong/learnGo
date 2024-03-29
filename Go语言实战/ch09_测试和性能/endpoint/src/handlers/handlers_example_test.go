// 9-26 handlers_example_test.go
// 这个示例程序展示如何编写基础示例
package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// ExampleSendJSON 提供了基础示例
func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)
	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}

	// 使用 fmt 将结果写到 stdout 来检测输出
	fmt.Println(u)
	// Output:
	// {Bill bill@ardanstudios.com}
}

// go test -v -run="ExampleSendJSON"

// G:\GitHub\learnGo\Go语言实战\ch09_测试和性能\endpoint>set GOPATH=G:\GitHub\learnGo\Go语言实战\ch09_测试和性能\endpoint
// G:\GitHub\learnGo\Go语言实战\ch09_测试和性能\endpoint\src\handlers>go test -v -run="ExampleSendJSON"
// === RUN   ExampleSendJSON
// --- PASS: ExampleSendJSON (0.00s)
// PASS
// ok      handlers        1.867s
