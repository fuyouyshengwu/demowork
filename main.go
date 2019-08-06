package main

import (
	"net/http"
	"os"
	"fmt"
)

func main(){
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		s := "<p>it is index</p>"
		writer.Write([]byte(s))
	})
	http.HandleFunc("/env", func(writer http.ResponseWriter, request *http.Request) {
		home := os.Getenv("JAVA_HOME")
		writer.Write([]byte(home))
	})
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("I am ok")
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("I am ok")
	})
	fmt.Println("working at 8080")
	http.ListenAndServe("0.0.0.0:8080",nil)
}