package main

import "fmt"

func main() {
    arr := [4]float64{2.5, 1.5, 3.0, 4.0}
    produto := 1.0
    for _, num := range arr {
        produto *= num
    }
    fmt.Println("Produto dos valores:", produto)
}
