package myinterface

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test2Slice(t *testing.T) {
	jsonData := `{"4": ["SOHOwireless-N", "SOHO250", "SOHO250wireless-N", "TZ 300"]}`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	value, ok := data["4"]
	if !ok {
		fmt.Println("Key not found")
		return
	}

	switch v := value.(type) {
	case []string: // 无法成功转换成 []string，但是可以断言成 []any
		fmt.Println("Value is a string slice", v)
	case []any:
		fmt.Println("Value is a string slice", v)
	default:
		fmt.Println("Unknown type")
	}
}
