package main

import (
        "fmt"
)

func main() {
        var s string = "hellok"
        reverse(s)
}

func reverse(s string) {
        r := []rune(s)
        length := len(r)
        for i := 0; i < length/2; i++ {
                r[i], r[length-i-1] = r[length-i-1], r[i]
        }
        s = string(r)
        fmt.Println(s)
}
