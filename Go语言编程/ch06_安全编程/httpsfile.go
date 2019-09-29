// 6-5 httpsfile
package main

import (
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("."))
	http.ListenAndServeTLS(":8001", "server.crt", "server.key", h)
	// http.ListenAndServe(":8001", h)
}
