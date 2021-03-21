package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":8080"
	if len(os.Args) >= 2 {
		addr = os.Args[1]
	}
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(addr, proxy))
}
