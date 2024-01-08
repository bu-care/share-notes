package main

import "fmt"

type A struct {
	name string
}

type B struct {
	model A
	id    string
}

func main() {
	a := A{name: "zhangsan"}
	b := B{id: "01", model: a}
	fmt.Println(b)
}
