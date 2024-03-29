// 9-19 handlers/handlers_test.go
// 这个示例程序展示如何测试内部服务端点
// 的执行效果
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"handlers"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	handlers.Routes()
}

// TestSendJSON 测试/sendjson 内部服务端点
func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.",
				ballotX, err)
		}
		t.Log("\tShould be able to create a request.",
			checkMark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}
		t.Log("\tShould decode the response.", checkMark)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name.", checkMark)
		} else {
			t.Error("\tShould have a Name.", ballotX, u.Name)
		}

		if u.Email == "bill@ardanstudios.com" {
			t.Log("\tShould have an Email.", checkMark)
		} else {
			t.Error("\tShould have an Email.", ballotX, u.Email)
		}
	}
}

// G:\GitHub\learnGo\Go语言实战\ch09_测试和性能\endpoint>set GOPATH=G:\GitHub\learnGo\Go语言实战\ch09_测试和性能\endpoint
// G:\GitHub\learnGo\Go语言实战\ch09_测试和性能\endpoint\src\handlers>go test handlers_test.go -v
// === RUN   TestSendJSON
// --- PASS: TestSendJSON (0.00s)
//         handlers_test.go:24: Given the need to test the SendJSON endpoint.
//         handlers_test.go:32:    Should be able to create a request. ✓
//         handlers_test.go:40:    Should receive "200" ✓
//         handlers_test.go:50:    Should decode the response. ✓
//         handlers_test.go:53:    Should have a Name. ✓
//         handlers_test.go:59:    Should have an Email. ✓
// PASS
// ok      command-line-arguments  2.097s
