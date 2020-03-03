package main

import (
    "fmt"   
)

func Main(req *Request) (interface{}, *Response) {
    fmt.Println("test")
    return "Hello World!", nil
}
