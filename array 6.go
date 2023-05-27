package main

import "fmt"

func main() {
    arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    var valor int
    fmt.Print("Informe um valor: ")
    fmt.Scanln(&valor)

    encontrado := false
    for _, num := range arr {
        if num == valor {
            encontrado = true
            break
        }
    }

    if encontrado {
        fmt.Println("O valor está presente no array.")
    } else {
        fmt.Println("O valor não está presente no array.")
    }
}
