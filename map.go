package main

import (
	"fmt"
	"reflect"
)

func main() {
	// panic: assignment to entry in nil map
	// 因為 m 是空的，還沒有指向任何指標。
	// var m map[string]int
	// m["age"] = 66

	m := make(map[string]int)
	// make function allocates and initialize a hash map data structure
	// and returns a map value that points to it
	m["age"] = 66
	m["name"] = 123
	fmt.Println(m)

	ks := reflect.ValueOf(m).MapKeys() // []Value

	fmt.Println(ks)

}
