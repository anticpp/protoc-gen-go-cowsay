package main

import (
	"fmt"
	"protoc-gen-go-cowsay/api/example"
)

func main() {
	m := example.Hello{Greeting: "Hello World!"}
	fmt.Println(m.Cowsay())
}
