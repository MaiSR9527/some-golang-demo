package main

import (
	"fmt"
	"net/http"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/17 10:07
 *@version v1.0
 */
func first(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		fmt.Fprintln(w, "header: "+key)
		for _, s := range value {
			fmt.Fprintln(w, "value: "+s)
		}
	}
	//get请求参数
	fmt.Fprintln(w, r.FormValue("name"))
	//post请求参数
	fmt.Fprintln(w, r.PostFormValue("name"))
	fmt.Fprintln(w, "firstFunction")
}

func main() {
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/first", first)
	server.ListenAndServe()
}
