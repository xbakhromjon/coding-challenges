package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	v := resty.Request{}
	fmt.Printf("%+v\n", v)
	fmt.Println("Hello World")
}
