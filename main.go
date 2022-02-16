package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(targetHost string) *httputil.ReverseProxy {
	url, err := url.Parse(targetHost)
	if err != nil {
		panic(err)
	}

	return httputil.NewSingleHostReverseProxy(url)
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	proxy := NewProxy("https://fplnow.com")

	http.HandleFunc("/", ProxyRequestHandler(proxy))
	http.ListenAndServe(":8080", nil)
}
