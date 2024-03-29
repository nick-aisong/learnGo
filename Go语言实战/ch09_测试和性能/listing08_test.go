// 9-8 listing08_test.go
// 这个示例程序展示如何写一个基本的表组测试
package listing08

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload 确认 http 包的 Get 函数可以下载内容
// 并正确处理不同的状态
func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.baidu.com",
			http.StatusOK,
		},
		{
			"https://www.sina.com/aaa",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
				u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.",
						ballotX, err)
				}
				t.Log("\t\tShould be able to Get the url",
					checkMark)

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v",
						u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status %v %v",
						u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}

// G:\GitHub\learnGo\Go语言实战\ch09_测试和性能>go test listing08_test.go -v
// === RUN   TestDownload
// --- PASS: TestDownload (0.19s)
//         listing08_test.go:30: Given the need to test downloading different content.
//         listing08_test.go:34:   When checking "http://www.baidu.com" for status code "200"
//         listing08_test.go:42:           Should be able to Get the url ✓
//         listing08_test.go:48:           Should have a "200" status. ✓
//         listing08_test.go:34:   When checking "https://www.sina.com/aaa" for status code "404"
//         listing08_test.go:42:           Should be able to Get the url ✓
//         listing08_test.go:48:           Should have a "404" status. ✓
// PASS
// ok      command-line-arguments  2.082s
