package main

import (
	"fmt"

	"go_code/chapter12/familyaccount-my/utils"
)

func main() {

	fmt.Println("这个是面向对象的方式完成~~")
	// 类似在python中，先初始化返回一个实例，然后用实例调用方法
	// utils.NewFamilyAccount().MainMenu() //思路非常清晰
	fa := utils.InitFamilyAccount()
	fa.Main_menu()
}
