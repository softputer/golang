package main

import (
	"sync"
	"fmt"
)

func main() {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	counter.m["key"] = 1
	counter.RLock()
	n := counter.m["key"]
	fmt.Println(n)
	counter.RUnlock()
	
	counter.Lock()
	counter.m["key"]++
	fmt.Println(counter.m["key"])
	counter.Unlock()	
}
