package main

import (
	"github.com/elazarl/goproxy"
	"net/http"
	"os"
	"strconv"
)

func main() {
	portStr := os.Getenv("PORT")
	var port int
	if portStr != "" {
		p, e := strconv.Atoi(portStr)
		if e != nil {
			return
		}
		port = p
	} else {
		port = 80
	}
	proxy := goproxy.NewProxyHttpServer()
	http.ListenAndServe(":"+strconv.Itoa(port), proxy)
}
