package main

import (
	"fmt"
	"net/http"
)

//web页面生成。
func testport(w http.ResponseWriter, r *http.Request) {
	for key, _ := range outputport {
		fmt.Fprintf(w, "%s\n", key)
	}
}

//web页面生成。
func testurl(w http.ResponseWriter, r *http.Request) {
	for key, _ := range outputurl {
		fmt.Fprintf(w, "%s\n", key)
	}
}
func webui() {
	server := http.Server{
		Addr: "0.0.0.0:8000",
	}
	http.HandleFunc("/testport", testport)
	http.HandleFunc("/testurl", testurl)
	server.ListenAndServe()
}
