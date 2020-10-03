package main

import "fmt"

var c, python, java bool

func printC() {
    c := true
    fmt.Println(c)
}

func main() {
    var i int
    printC()
    fmt.Println(i, c, python, java)
}
