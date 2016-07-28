package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name	string
	Body	string
	Time	int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	fmt.Println(b)

	var result Message
	err := json.Unmarshal(b, &result)
	fmt.Println(result)
	fmt.Println(err)
	if err != nil {
		fmt.Println(result)
	}

	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777
	
	switch v := i.(type) {
	case int:
		fmt.Println("int", v*2)
	case float64:
		fmt.Println(v)
		fmt.Println("float", 1/v)
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}
	r := i.(float64)
	fmt.Println("the circle's area", r)
}
