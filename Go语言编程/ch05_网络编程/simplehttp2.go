// 5-3 simplehttp2.go
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//G:\GitHub\learnGo\Go语言编程\ch05_网络编程>go run simplehttp2.go www.sina.com:80
//HTTP/1.1 404 Not Found on Accelerator
//Server: nginx
//Date: Sat, 28 Sep 2019 10:40:31 GMT
//Content-Type: text/html
//Content-Length: 297
//Connection: close
//Via: http/1.1 cnc.ningbo.ha2ts4.25 (ApacheTrafficServer/6.2.1 [c s f ])
//Cache-Control: no-store
//Content-Language: en
//X-Cache: MISS.25
//X-Via-CDN: f=edge,s=cnc.ningbo.ha2ts4.21.nb.sinaedge.com,c=112.64.61.38;f=Edge,s=cnc.ningbo.ha2ts4.25,c=101.71.100.21
//X-Via-Edge: 1569667231011263d40707b6447653cc908f7
