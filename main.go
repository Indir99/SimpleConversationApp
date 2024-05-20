package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func main() {
	var i = f()
	var j = f()
	fmt.Printf("Address of i = %v, address of j = %v \n", i, j)
}
