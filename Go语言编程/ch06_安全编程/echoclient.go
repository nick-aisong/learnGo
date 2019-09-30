// 6-7 echoclient
package main

import (
	"crypto/tls"
	"io"
	"log"
)

func main() {
	// 如果config 参数是nil，会报错：dial: x509: certificate signed by unknown authority
	// conn, err := tls.Dial("tcp", "127.0.0.1:8000", nil)
	// 忽略CA是否可信任校验，单向信任服务端证书
	conn, err := tls.Dial("tcp", "127.0.0.1:8000", &tls.Config{RootCAs: nil, InsecureSkipVerify: true})
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())
	state := conn.ConnectionState()
	log.Println("client: handshake: ", state.HandshakeComplete)
	log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)
	message := "Hello\n"
	n, err := io.WriteString(conn, message)
	if err != nil {
		log.Fatalf("client: write: %s", err)
	}
	log.Printf("client: wrote %q (%d bytes)", message, n)
	reply := make([]byte, 256)
	n, err = conn.Read(reply)
	log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
	log.Print("client: exiting")
}

//G:\GitHub\learnGo\Go语言编程\ch06_安全编程>go run echoclient.go
//2019/09/29 22:19:36 client: dial: x509: certificate signed by unknown authority
//exit status 1

//G:\GitHub\learnGo\Go语言编程\ch06_安全编程>go run echoclient.go
//2019/09/29 22:24:53 client: connected to:  127.0.0.1:8000
//2019/09/29 22:24:53 client: handshake:  true
//2019/09/29 22:24:53 client: mutual:  true
//2019/09/29 22:24:53 client: wrote "Hello\n" (6 bytes)
//2019/09/29 22:24:53 client: read "Hello\n" (6 bytes)
//2019/09/29 22:24:53 client: exiting
