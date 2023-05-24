package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	log.Printf("Client IP: %s, HTTP StatusCode: %d", r.RemoteAddr, http.StatusOK)

	// 4.处理 "/healthz" 路径，返回200
	if r.URL.Path == "/healthz" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	for name, values := range r.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// 2.读取系统环境变量中的VERSION配置
	version := os.Getenv("GOPATH")
	if version != "" {
		// 将VERSION配置写入响应的header
		w.Header().Set("GOPATH", version)
	}

	io.WriteString(w, "ok\n")
}
