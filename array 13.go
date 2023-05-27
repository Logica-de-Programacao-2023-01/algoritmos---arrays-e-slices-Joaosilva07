package main

import "fmt"

func main() {
    arr := [7]int{1, 2, 3, 4, 5, 6, 7}

    var numero int
    fmt.Print("Informe um número: ")
    fmt.Scanln(&numero)

    arr[0] += numero
    arr[len(arr)-1] += numero

    fmt.Println("Array resultante:", arr)
}
