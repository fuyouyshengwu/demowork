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
	fmt.Println("working at 8080")
	http.ListenAndServe("127.0.0.1:8080",nil)
}