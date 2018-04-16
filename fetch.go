package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix("https://", url) {
			url = "https://" + url
		}

		resp, err := http.Get(url)

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}

		resp.Body.Close()
		fmt.Fprintf(os.Stdout, "%s", resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
