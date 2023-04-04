package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello")
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})

	http.HandleFunc("/hell", func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Fprintf(writer, "hello world")
		for s := range request.Header {
			fmt.Fprintf(writer, "%+v", s)
		}
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})

	http.ListenAndServe(":9999", nil)
}
