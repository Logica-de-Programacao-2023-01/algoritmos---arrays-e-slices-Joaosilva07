package main

import "fmt"

func main() {
    slice := []string{"apple", "banana", "orange", "banana", "grape", "banana", "kiwi", "banana"}

    var valor string
    fmt.Print("Informe um valor: ")
    fmt.Scanln(&valor)

    novoSlice := []string{}
    for _, str := range slice {
        if str != valor {
            novoSlice = append(novoSlice, str)
        }
    }

    fmt.Println("Slice resultante:", novoSlice)
}
