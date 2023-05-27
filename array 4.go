package main

import (
    "fmt"
)

func main() {
    var tamanho int
    fmt.Print("Informe o tamanho do slice: ")
    fmt.Scanln(&tamanho)

    slice := make([]int, tamanho)

    for i := 0; i < tamanho; i++ {
        fmt.Printf("Informe o valor do elemento %d: ", i+1)
        fmt.Scanln(&slice[i])
    }

    fmt.Println("Slice:", slice)

    soma := 0
    for _, num := range slice {
        soma += num
    }
    fmt.Println("Soma dos valores:", soma)
}
