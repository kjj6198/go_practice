package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:3001", nil))
	
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%s %s %s\n", req.Method, req.URL, req.Proto)

	for key, value := range req.Header {
		fmt.Fprintf(res, "Header[%q] = %q\n", key, value)
	}
	fmt.Fprintf(res, "Host = %q\n", req.Host)
	fmt.Fprintf(res, "URL.Path = %q \n", req.URL.Path)

	if err:= req.ParseForm(); err != nil {
		log.Print(err)
	}

	for key, value := range req.Form {
		fmt.Fprintf(res, "Form[%q] = %q\n", key, value)
	}
}