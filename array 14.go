package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

    var indice1, indice2 int
    fmt.Print("Informe o primeiro índice: ")
    fmt.Scanln(&indice1)
    fmt.Print("Informe o segundo índice: ")
    fmt.Scanln(&indice2)

    slice[indice1], slice[indice2] = slice[indice2], slice[indice1]

    fmt.Println("Slice resultante:", slice)
}
