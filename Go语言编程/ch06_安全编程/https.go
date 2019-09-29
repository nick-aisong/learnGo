// 6-3
package main

import (
	"fmt"
	"net/http"
)

const SERVER_PORT = 8080
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "hello"

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}
func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	// HTTPS
	http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "server.crt", "server.key", nil)

	// HTTP
	// http.ListenAndServe(fmt.Sprintf(":%d", SERVER_PORT), nil)
}
