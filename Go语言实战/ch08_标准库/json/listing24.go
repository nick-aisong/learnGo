// 8-24 listing24.go
// 这个示例程序展示如何使用 json 包和 NewDecoder 函数
// 来解码 JSON 响应
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// gResult 映射到从搜索拿到的结果文档
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	// gResponse 包含顶级的文档
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// 向 Google 发起搜索
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	// 将 JSON 响应解码到结构类型
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}

// 8-25 golang.org/src/encoding/json/stream.go
// NewDecoder 返回从 r 读取的解码器
//
// 解码器自己会进行缓冲，而且可能会从 r 读比解码 JSON 值
// 所需的更多的数据
// func NewDecoder(r io.Reader) *Decoder
// Decode 从自己的输入里读取下一个编码好的 JSON 值，
// 并存入 v 所指向的值里
//
// 要知道从 JSON 转换为 Go 的值的细节，
// 请查看 Unmarshal 的文档
// func (dec *Decoder) Decode(v interface{}) error
