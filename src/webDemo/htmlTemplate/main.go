package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/17 10:28
 *@version v1.0
 */
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
}
func main() {
	server := http.Server{Addr: "localhost:8080"}
	//http://localhost:8080/static/js/index.js  js/index.js转发至/static/文件夹 使文件服务生效
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
