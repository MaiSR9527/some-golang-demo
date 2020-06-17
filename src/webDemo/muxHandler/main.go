package main

import (
	"fmt"
	"net/http"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/17 09:56
 *@version v1.0
 */
/*type HandlerOne struct {
}
type HandlerTwo struct {
}
type HandlerThree struct {
}

func (m *HandlerOne) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HandlerOne"))
}
func (m *HandlerTwo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HandlerTwo"))
}
func (m *HandlerThree) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HandlerThree"))
}

func main() {
	//定义多个结构体，并重写接口
	h1 := HandlerOne{}
	h2 := HandlerTwo{}
	h3 := HandlerThree{}
	server := http.Server{Addr: "localhost:8080"}
	http.Handle("/one", &h1)
	http.Handle("/two", &h2)
	http.Handle("/three", &h3)
	server.ListenAndServe()

}*/

//多处理函数方式
func first(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "firstFunction")
}
func second(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "firstFunction")
}
func third(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "firstFunction")
}

func main() {
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/first", first)
	http.HandleFunc("/second", second)
	http.HandleFunc("/third", third)
	server.ListenAndServe()
}
