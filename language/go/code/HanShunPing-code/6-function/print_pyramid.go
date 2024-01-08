package main

import "fmt"

func printPyramid(totalLevel int){
	for i := 1; i <= totalLevel; i++ {
		//在打印*前先打印空格
		for k := 1; k <= totalLevel - i; k++ {
			fmt.Print(" ") // Print函数结尾不会换行，Println打印结尾会换行
		}

		//j 表示每层打印多少*
		for j :=1; j <= 2 * i - 1; j++ {
				fmt.Print("*")
		}
		fmt.Println()
	}
}


func main() {
	var num int
	fmt.Println("please enter a number...")
	fmt.Scanf("%d", &num)
	printPyramid(num)
}