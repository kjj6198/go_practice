package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  counts := make(map[string]int) // create a map
  files := os.Args[1:]
  if len(files) == 0 {
  	countLines(os.Stdin, counts)
  } else {
  	for _, arg := range files {
  		f, err := os.Open(arg)
  		if err != nil {
  			// print error message and continue.
  			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
  			continue

  		}

  		countLines(f, counts)
  		f.Close()
   	}
  }

  for key, value := range counts {
  	if value > 1 {
  		fmt.Printf("%d\t%s\n", value, key)
  		
  	}
  }
}

func countLines(f *os.File, counts map[string]int) {
  input := bufio.NewScanner(f)

  for input.Scan() {
	  counts[input.Text()]++
  }
}