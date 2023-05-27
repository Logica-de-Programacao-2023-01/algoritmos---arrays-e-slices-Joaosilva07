package main

import "fmt"

func main() {
    arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}

    ordenado := true
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] > arr[i+1] {
            ordenado = false
            break
        }
    }

    if ordenado {
        fmt.Println("O array está ordenado em ordem crescente.")
    } else {
        fmt.Println("O array não está ordenado em ordem crescente.")
    }
}
