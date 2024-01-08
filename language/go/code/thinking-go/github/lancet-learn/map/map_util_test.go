package map_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/strutil"
	// "github.com/duke-git/lancet/v2/maputil"
)

func TestMap(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City bool   `json:"city"`
	}

	// var p Person
	// maputil.MapTo(map2, &p)
	p := &Person{
		Name: "xbu",
		Age:  20,
		City: true,
	}
	fmt.Println(p) // &{xbu 20 true}

	map2 := map[string]any{"name": nil, "age": 35, "city": false}
	data, _ := json.Marshal(map2)
	// nil放进 Person 结构体中不会改变原本的值
	json.Unmarshal(data, p)
	fmt.Println(p) // &{xbu 35 false}

	s := "hello"
	rs := strutil.Reverse(s)
	fmt.Println(rs) //olleh
}
