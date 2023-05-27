package main

import "fmt"

func main() {
    arr := [3]int{2, 4, 6}
    soma := 0
    for _, num := range arr {
        soma += num
    }
    fmt.Println("Soma dos valores:", soma)
}
