package main

import "fmt"

func main() {
    arr := [10]float64{2.3, 5.6, 8.1, 4.2, 6.7, 9.4, 3.5, 7.8, 1.9, 0.5}

    novoSlice := []float64{}
    for _, num := range arr {
        if num > 5 {
            novoSlice = append(novoSlice, num)
        }
    }

    fmt.Println("Novo Slice:", novoSlice)
}
