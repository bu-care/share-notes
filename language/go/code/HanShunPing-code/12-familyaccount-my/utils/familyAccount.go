package utils

import (
	"fmt"
)

type FamilyAccount struct {
	username string
	password string

	//声明一个变量，保存接收用户输入的选项
	key string
	//声明一个变量，控制是否退出for
	loop bool

	//定义账户的余额 []
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义个变量，记录是否有收支的行为
	flag bool
	//收支的详情使用字符串来记录
	//当有收支时，只需要对details 进行拼接处理即可
	details          string
	transfer_account string
}

//类似于python中的__init__(self)，初始化实例属性
func InitFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		username:         "admin",
		password:         "password",
		key:              "",
		loop:             true,
		balance:          10000.0,
		money:            0.0,
		note:             "",
		flag:             false,
		details:          "收支\t账户金额\t收支金额\t说明",
		transfer_account: "",
	}
}

func (fa *FamilyAccount) show_details() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if fa.flag {
		fmt.Println(fa.details)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}

func (fa *FamilyAccount) income() {
	fmt.Println("plese input revenue amount:")
	fmt.Scanln(&fa.money)
	fa.balance += fa.money
	fmt.Println("input income note:")
	fmt.Scanln(&fa.note)
	fa.details += fmt.Sprintf("\ndetails\t%v\t%v\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

func (fa *FamilyAccount) pay() {
	fmt.Println("plese input pay amount:")
loop:
	fmt.Scanln(&fa.money)
	if fa.money > fa.balance {
		fmt.Println("pay amount more than balance, please re-enter...")
		goto loop
	}
	fa.balance -= fa.money
	fmt.Println("input pay note:")
	fmt.Scanln(&fa.note)
	fa.details += fmt.Sprintf("\ndetails\t%v\t%v\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

func (fa *FamilyAccount) transfer() {
	fmt.Println("plese input transfer account:")
	fmt.Scanln(&fa.transfer_account)
	fmt.Println("plese input transfer amount:")
loop:
	fmt.Scanln(&fa.money)
	if fa.money > fa.balance {
		fmt.Println("transfer amount more than balance, please re-enter...")
		goto loop
	}
	fa.balance -= fa.money
	fmt.Println("input transfer note:")
	fmt.Scanln(&fa.note)
	fa.note += "[transfer account: " + fa.transfer_account + "]"
	fa.details += fmt.Sprintf("\ndetails\t%v\t%v\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

func (fa *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}

	if choice == "y" {
		fa.loop = false
	}
}

func (fa *FamilyAccount) Main_menu() {
	count := 0
	var username string
	var password string
	for {
		count += 1
		fmt.Print("please input username: ")
		fmt.Scanln(&username)
		fmt.Print("please input password: ")
		fmt.Scanln(&password)
		if username == fa.username && password == fa.password {
			fmt.Println("login success")
			break
		} else {
			fmt.Println("username or password incorrect, please re-enter...")
		}
		if count >= 3 {
			fmt.Println("username or password error 3 times, end diag!")
			return
		}

	}
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Println("                  5 转账")
		fmt.Print("请选择(1-5): ")
		fmt.Scanln(&fa.key)

		switch fa.key {
		case "1":
			fa.show_details()
		case "2":
			fa.income()
		case "3":
			fa.pay()
		case "4":
			fa.exit()
		case "5":
			fa.transfer()
		default:
			fmt.Println("Unknown key: ", fa.key)
		}
		if !fa.loop {
			break
		}
	}
}

//在utils中写main是没有用的，在go中，程序需要从main包开始执行，无法从当前脚本执行
func main() {
	fmt.Println("Starting...")
	fa := InitFamilyAccount()
	fa.Main_menu()
}
