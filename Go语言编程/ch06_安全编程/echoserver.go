// 6-6 echoserver
package main

import (
	"crypto/rand"
	"crypto/tls"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = time.Now
	config.Rand = rand.Reader
	service := "127.0.0.1:8000"
	listener, err := tls.Listen("tcp", service, &config)

	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	for {
		log.Print("server: conn: waiting")
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("server: conn: read: %s", err)
			}
			break
		}
		log.Printf("server: conn: echo %q\n", string(buf[:n]))
		n, err = conn.Write(buf[:n])
		log.Printf("server: conn: wrote %d bytes", n)
		if err != nil {
			log.Printf("server: write: %s", err)
			break
		}
	}
	log.Println("server: conn: closed")
}

//G:\GitHub\learnGo\Go语言编程\ch06_安全编程>go run echoserver.go
//2019/09/29 22:17:13 server: listening
//2019/09/29 22:19:36 server: accepted from 127.0.0.1:37493
//2019/09/29 22:19:36 server: conn: waiting
//2019/09/29 22:19:36 server: conn: read: remote error: tls: bad certificate
//2019/09/29 22:19:36 server: conn: closed

//2019/09/29 22:24:53 server: accepted from 127.0.0.1:37669
//2019/09/29 22:24:53 server: conn: waiting
//2019/09/29 22:24:53 server: conn: echo "Hello\n"
//2019/09/29 22:24:53 server: conn: wrote 6 bytes
//2019/09/29 22:24:53 server: conn: waiting
//2019/09/29 22:24:53 server: conn: closed
