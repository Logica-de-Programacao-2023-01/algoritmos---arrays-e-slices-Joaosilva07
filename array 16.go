package main

import "fmt"

func main() {
    arr := [10]int{2, 4, 6, 7, 9, 10, 12, 15, 17, 18}

    novoSlice := []int{}
    for _, num := range arr {
        if num%2 == 0 {
            novoSlice = append(novoSlice, num)
        }
    }

    fmt.Println("Novo Slice:", novoSlice)
}
