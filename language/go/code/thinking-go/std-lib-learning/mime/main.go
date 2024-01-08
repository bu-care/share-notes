package main

import (
	"fmt"
	"mime"
)

func main() {
	mineType1 := mime.TypeByExtension(".png")
	fmt.Println("mineType1: ", mineType1)
}
