package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/18 11:48
 *@version v1.0
 */
func index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintln(w, vars["url"])

}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sxt/{url}", index)
	http.ListenAndServe("localhost:8089", router)
}
