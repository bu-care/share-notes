package ch06

import (
	"fmt"
	"testing"
)

func TestIntSet(t *testing.T) {
	var x IntSet
	x.Add(10)
	x.Add(9)
	x.Remove(8)
	x.Add(112)
	// x.Add(113)
	// x.Remove(112)
	// x.Clear()
	fmt.Println(x.String())
	fmt.Println(x.Len())
	y := x.Copy()
	y.AddAll(1, 2, 3)
	fmt.Println(y.String())

	// s := fmt.Sprintf("%d", ^uint(0))
	// b := 32 << (^uint(0) >> 63)
	// fmt.Println(s, b)
	// fmt.Printf("%b, %b, %d", uint(0), ^uint(0), len(s))
	// fmt.Println("elems: ", y.Elems())
}
