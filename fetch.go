package main

import (
  "fmt"
  "net/http"
  "os"
  "strings"
	"io"
	"log"
)

func main() {
	url := "buzzorange.com/"

	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	res, err1 := http.Get(url)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err1)
		os.Exit(1)
	}

	_, err2 := io.Copy(os.Stdout, res.Body)
	res.Body.Close()

	if err2 != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err2)
		os.Exit(1)
	}

	// fmt.Printf("%s", data)
	fmt.Printf("%s", res.Status)

	r := strings.NewReader("some io.reader Stream to be read")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
		res.StatusCode
	}
}