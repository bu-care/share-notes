package map_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/yudai/gojsondiff"
)

func TestProcess(t *testing.T) {
	// a := map[string]any{
	// 	"a": "0",
	// }
	// c := map[string]any{
	// 	"b": a["b"],
	// }
	// fmt.Println(c)

	// json1 := `{"name": "John", "age": 30, "city": "New York"}`
	// json2 := `{"name": "John", "age": 35, "city": "Boston"}`

	// map1 := map[string]any{"name": "John", "age": 30, "city": "New York"}
	var map1 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City bool   `json:"city"`
	}
	map1.Name = "John"
	map1.Age = 30
	// map1.City = "New York"
	map2 := map[string]any{"name": "John", "age": 35, "city": nil}
	json1, _ := json.Marshal(map1)
	json2, _ := json.Marshal(map2)

	diff, err := gojsondiff.New().Compare([]byte(json1), []byte(json2))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(diff.Deltas(), diff.Modified())

	map_str := fmt.Sprint(map2)
	fmt.Println(map_str, string(json2))
}
