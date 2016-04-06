package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main() {
	url := "http://buzzorange.com/"
	res, err1 := http.Get(url)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err1)
		os.Exit(1)
	}

	b, err2 := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err2 != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err2)
		os.Exit(1)
	}

	fmt.Printf("%s", b)
}