package main

import "fmt"

func main() {
    arr := [5]int{9, 2, 7, 6, 5}

    novoSlice := []int{}
    for _, num := range arr {
        if num%3 == 0 {
            novoSlice = append(novoSlice, num)
        }
    }

    fmt.Println("Novo Slice:", novoSlice)
}
