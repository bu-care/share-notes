package main

import (
	"fmt"
	"go_code/chapter15/unittest/monster"
)

func main() {
	m := monster.InitMonster()
	m.Store()
	fmt.Println("---------")
	m.Restore()
}
