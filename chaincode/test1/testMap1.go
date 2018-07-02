package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//m1 := map[string]interface{}{"name": "John", "age": 10}

	m1 := make(map[string]string)
	m1["name"] = "John"
	m1["age"] = "10"


	str, err := json.Marshal(m1)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	fmt.Println("str:", string(str))

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		fmt.Println(dat)
		fmt.Println(dat["name"])
	} else {
		fmt.Println(err)
	}

	// 查找键值是否存在
	if v, ok := dat["name"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	// 遍历map
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}