package main

import (
  "fmt"
  "os"
  "path"
)

func main() {
	// 如果一開始沒有賦值，會被隱式賦值
	var result string
	var resultWithName string
	var sep string	
	// 等價
	// s := ""
  

	// go 裡面只允許 i++ 但不允許 ++i，也不允許賦值。
		
	for i:= 1; i < len(os.Args); i++ {
	  result += sep + os.Args[i]
	  sep = " "
	}

	fmt.Println(result)


	for i:= 0; i < len(os.Args); i++ {
	  resultWithName += sep + os.Args[i]
	  sep = "\n"
	}

	// 1.1
	// 1.2
	fmt.Println(path.Base(resultWithName))
	
}