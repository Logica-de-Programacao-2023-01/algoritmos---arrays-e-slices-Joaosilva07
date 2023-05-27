package main

import "fmt"

func main() {
    var n int
    fmt.Print("Informe o tamanho dos arrays: ")
    fmt.Scanln(&n)

    arr1 := make([]int, n)
    arr2 := make([]int, n)

    fmt.Println("Informe os elementos do primeiro array:")
    for i := 0; i < n; i++ {
        fmt.Printf("Elemento %d: ", i+1)
        fmt.Scanln(&arr1[i])
    }

    fmt.Println("Informe os elementos do segundo array:")
    for i := 0; i < n; i++ {
        fmt.Printf("Elemento %d: ", i+1)
        fmt.Scanln(&arr2[i])
    }

    arr3 := make([]int, n)
    for i := 0; i < n; i++ {
        arr3[i] = arr1[i] + arr2[i]
    }

    fmt.Println("Terceiro array (soma):", arr3)
}
