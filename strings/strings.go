package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {
    str := "Hello World! HHH"
    p("Replace 1: ", s.Replace(str, "H", "h", 1))
    p("Repalce 3: ", s.Replace(str, "H", "h", 3))
    p("Replace -1: ", s.Replace(str, "H", "h", -1))
    p("Compare: ", s.Compare(str, "J"))
    p("Contains: ", s.Contains(str, "HH"))
    p("ContainsAny: ", s.ContainsAny(str, ""))
    p("ContainsRune: ", s.ContainsRune(str, 33))
    p("Count: ", s.Count(str, "HH"))
}
