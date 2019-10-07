// 9-12 listing12_test.go
// 这个示例程序展示如何内部模仿 HTTP GET 调用
// 与本书之前的例子有些差别
package listing12

import (
	//"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// feed 模仿了我们期望接收的 XML 文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
	<channel>
		<title>Going Go Programming</title>
		<description>Golang : https://github.com/goinggo</description>
			<link>http://www.goinggo.net/</link>
		<item>
			<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
			<title>Object Oriented Programming Mechanics</title>
			<description>Go is an object oriented language.</description>
			<link>http://www.goinggo.net/2015/03/object-oriented</link>
		</item>
	</channel>
</rss>`

// mockServer 返回用来处理请求的服务器的指针
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload 确认 http 包的 Get 函数可以下载内容
// 并且内容可以被正确地反序列化并关闭
func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v",
				statusCode, checkMark)
		}
	}
}

// G:\GitHub\learnGo\Go语言实战\ch09_测试和性能>go test listing12_test.go -v
// === RUN   TestDownload
// --- PASS: TestDownload (0.01s)
//         listing12_test.go:52: Given the need to test downloading content.
//         listing12_test.go:55:   When checking "http://127.0.0.1:27051" for status code "200"
//         listing12_test.go:63:           Should be able to make the Get call. ✓
//         listing12_test.go:72:           Should receive a "200" status. ✓
// PASS
// ok      command-line-arguments  1.812s
