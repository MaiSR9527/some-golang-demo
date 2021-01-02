package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/**
 *
 *@author MaiShuRen
 *@date 2021/1/2 15:44
 *@version v1.0
 */

const Addr = ":9088"

func main() {
	// 创建路由器
	mux := http.NewServeMux()
	// 设置路由规则
	mux.HandleFunc("/bye", sayBye)
	// 创建服务器
	server := &http.Server{
		Addr:         Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 3,
	}
	log.Println("Start http server at " + Addr)
	log.Fatal(server.ListenAndServe())
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	i, err := w.Write([]byte("bye bye http server"))

	fmt.Println(i, err)
}
