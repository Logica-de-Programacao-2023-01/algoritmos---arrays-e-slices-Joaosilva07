package main

import "fmt"

func main() {
    arr := [6]float64{1.2, 2.4, 3.6, 4.8, 6.0, 7.2}

    var numero float64
    fmt.Print("Informe um número: ")
    fmt.Scanln(&numero)

    for i := 0; i < len(arr); i++ {
        arr[i] += numero
    }

    fmt.Println("Array resultante:", arr)
}
