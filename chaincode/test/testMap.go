package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	map1 := make(map[string]interface{})
	map1["1"] = "hello"
	map1["2"] = "world"
	//return []byte
	str, err := json.Marshal(map1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("map to json", string(str))

	//json([]byte) to map
	map2 := make(map[string]interface{})
	err = json.Unmarshal(str, &map2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json to map ", map2)
	fmt.Println("The value of key1 is", map2["1"])

}