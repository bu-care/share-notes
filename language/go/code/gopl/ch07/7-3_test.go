package ch07

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	fmt.Println(data)
	Sort(data)
	fmt.Println("after sort: ", data)
}

func TestTree(t *testing.T) {
	//初始：nil<-2->nil
	tree := &tree{value: 9}
	//加入8后：nil<-2->8->nil
	// add(tree, 8)
	// add(tree, 15)
	// add(tree, 5)
	addAll(tree, []int{8, 15, 5})
	fmt.Println(tree)
}
