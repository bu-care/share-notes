package main

import "fmt"

func main() {
	// 9的ASCII码是57
	a := ['9']int32{
		'0': '1',
		'1': '2',
	}
	fmt.Println(a, len(a), '9')

	// 上面的写法等价于这样，切片第48的位置赋值为49（b[48] = 49）
	b := [57]int32{
		48: 49,
		49: 50,
	}
	fmt.Println(b, len(b))
}
