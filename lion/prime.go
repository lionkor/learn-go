package main

import "fmt"

func isPrime(x int) bool {
    for i := 2; i < x; i++ {
        if x % i == 0 {
            return false
        }
    }
    return true
}

func main() {
    for i := 1; i < 1000; i++ {
        res := isPrime(i)
        if res {
            fmt.Print(i, " ")
        }
    }
    fmt.Print("\n")
}
