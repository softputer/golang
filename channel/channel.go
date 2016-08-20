package main

import (
	"fmt"
)

func pass(ch chan<- int) {
	ch <- 3
}

func main() {
	fmt.Println("Helo")
	ch := make(chan int, 1)
	go pass(ch)
	mes := <- ch
	fmt.Println(mes)
}
