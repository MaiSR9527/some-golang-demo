package main

import (
	"fmt"
	"net/http"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/17 09:31
 *@version v1.0
 */
func main() {
	http.HandleFunc("/", welcome)
	http.ListenAndServe("localhost:8080", nil)
}
func welcome(res http.ResponseWriter, rquest *http.Request) {
	res.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(res, "<a href='https://www.baidu.com' target='_target'>baidu</a>")
}
