package main

import "net/http"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/17 09:45
 *@version v1.0
 */
type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MyHandler"))
}

func main() {
	h := MyHandler{}
	//处理所有请求，/ /**等都会被路由到MyHandler
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &h,
	}
	server.ListenAndServe()
}
