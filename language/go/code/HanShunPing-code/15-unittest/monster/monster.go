package monster

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func InitMonster() *Monster {
	return &Monster{
		Name:  "abc",
		Age:   22,
		Skill: "catch fish",
	}
}

func (m *Monster) Store() bool {
	arr_data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error marshallingMonster: ", err)
		return false
	}
	data := string(arr_data)
	fmt.Println("Marshaling data: ", data)

	file_path := "./monster.json"
	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error marshallingMonster: ", err)
		return false
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(data)
	writer.Flush()
	return true
}

func (m *Monster) Restore() bool {
	file_path := "./monster.json"
	data, err := ioutil.ReadFile(file_path) //返回的data是 byte 数组
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return false
	}
	// fmt.Println("json data: ", data)
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
		return false
	}
	fmt.Println("Monster Unmarshal: ", *m)
	return true
}
