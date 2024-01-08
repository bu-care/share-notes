package myinterface

import (
	"fmt"
	"testing"
)

type NetworksIcon struct {
	File_name   string `bson:"file_name" json:"file_name"`
	S3_file_key string `bson:"s3_file_key" json:"s3_file_key"`
}

type Networks struct {
	Name         string       `bson:"name" json:"name"`
	Icon_setting NetworksIcon `bson:"icon_setting" json:"icon_setting"`
}

func interface2int(input any) any {
	var res any
	fmt.Printf("input type: %t", input)
	// res_j := input.(primitive.DateTime)
	// res = int(res_j)
	switch res := input.(type) {
	case int: // 类型断言，断言成功后结果 res 就转换为该类型
		fmt.Println("----", res)
	case Networks:
		// res = input.(Networks)
		fmt.Println("----", res)
	default:
		// res = input
		fmt.Println("+++++", res)
	}
	return res
}

func Test2Int(t *testing.T) {
	input := `{
		"created_time": 1676276492984,
		"icon": {
			"file_id": null,
			"file_name": "test.png",
			"url": null
		},
		"id": "63e9f30c26321993c62abdff",
		"name": "go test"
	}`
	res := interface2int(input)
	fmt.Println("result:", res)
}
