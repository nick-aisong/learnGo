// 9-1 listing01_test.go
// 这个示例程序展示如何写基础单元测试
package listing01

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload 确认 http 包的 Get 函数可以下载内容
func TestDownload(t *testing.T) {
	url := "http://www.baidu.com"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v",
					statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}

// G:\GitHub\learnGo\Go语言实战\ch09_测试和性能>go test listing01_test.go -v
// === RUN   TestDownload
// --- PASS: TestDownload (0.04s)
//         listing01_test.go:18: Given the need to test downloading content.
//         listing01_test.go:21:   When checking "http://www.baidu.com" for status code "200"
//         listing01_test.go:29:           Should be able to make the Get call. ✓
//         listing01_test.go:35:           Should receive a "200" status. ✓
// PASS
// ok      command-line-arguments  1.908s
