package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	helloOtus := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(helloOtus))
}
