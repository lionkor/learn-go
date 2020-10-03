package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    fmt.Println("my favourite person is YOU and my favourite number is not", rand.Intn(100))
}
