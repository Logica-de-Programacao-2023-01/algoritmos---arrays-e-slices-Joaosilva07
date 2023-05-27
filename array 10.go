package main

import (
    "fmt"
    "math"
)

func main() {
    slice := []int{3, 5, 1, 2, 4, 7, 9, 6, 8, 0}

    min := math.MaxInt64
    max := math.MinInt64

    for _, num := range slice {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }

    fmt.Println("Valor mínimo:", min)
    fmt.Println("Valor máximo:", max)
}
