package main

import (
	"net/http"
	"os"
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
	http.ListenAndServe("127.0.0.1:80",nil)
}