package main

import (
	"fmt"
)

func main() {
	//Create Map
	m := make(map[string]string)
	m["network"] = "network"
	fmt.Println(m["network"])
	//Check if network is in map
	if _, ok := m["network"]; ok {
		fmt.Println("network is in map")
	}
}
