package main

import (
	"encoding/json"
	"fmt"

	"github.com/creasty/defaults"
)

type Member struct {
	Name string `json:"name"`
}

type modelType interface {
	Member
}

type ModelOperate[T any] struct {
	Id    int `bson:"_id" json:"id"`
	Model T   `bson:",inline"`
}

func interface2map(v interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println("interface Marshal error: ", err)
		return nil, err
	}
	var map_json map[string]interface{}
	err = json.Unmarshal(data, &map_json)
	if err != nil {
		fmt.Println("Unmarshal to map error: ", err)
		return nil, err
	}

	return map_json, nil
}

func map2modelInstance[T any](input_map any) (*T, error) {
	data, err := json.Marshal(input_map)
	if err != nil {
		fmt.Println("interface Marshal error: ", err)
		return nil, err
	}
	model_instance := new(T)
	err = json.Unmarshal([]byte(data), model_instance)
	if err != nil {
		fmt.Println("Unmarshal to map error: ", err)
		return nil, err
	}
	// fmt.Printf("model_instance: %p---%v\n", model_instance, model_instance)

	return model_instance, nil
}

func Operator[T any](inputInstance T) ModelOperate[T] {
	var model_instance *T
	//using "default" tag to set default struct values
	default_instance := new(T)
	err := defaults.Set(default_instance)
	if err != nil {
		fmt.Println("set default model err: ", err)
	} else {
		init_map, err := interface2map(new(T))
		if err != nil {
			fmt.Println("init_instance err: ", err)
			goto operate_model_tag
		}
		input_map, err := interface2map(inputInstance)
		if err != nil {
			fmt.Println("inputInstance err: ", err)
			goto operate_model_tag
		}
		default_map, err := interface2map(default_instance)
		if err != nil {
			fmt.Println("default_instance err: ", err)
			goto operate_model_tag
		}
		// cur_time := primitive.NewDateTimeFromTime(time.Now().UTC())
		for k, v := range default_map {
			if default_map[k] != init_map[k] {
				if input_map[k] == init_map[k] {
					input_map[k] = v
				}
			}
			// if slices.Contains([]string{"created_time", "created_at"}, k) || (k == "modified_time") {
			// 	input_map[k] = cur_time
			// }
		}
		fmt.Println("input_map: ", input_map)
		model_instance, _ = map2modelInstance[T](input_map)
		// fmt.Printf("model_instance: %p---%v\n", model_instance, model_instance)
	}
operate_model_tag:
	if model_instance == nil {
		model_instance = &inputInstance
	}
	operateModel := ModelOperate[T]{
		Model: *model_instance,
	}
	return operateModel
}
