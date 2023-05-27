package main

import "fmt"

func main() {
    var n int
    fmt.Print("Informe um número inteiro positivo: ")
    fmt.Scanln(&n)

    numerosPrimos := []int{}
    i := 2

    for len(numerosPrimos) < n {
        if isPrimo(i) {
            numerosPrimos = append(numerosPrimos, i)
        }
        i++
    }

    fmt.Println("Array de números primos:", numerosPrimos)
}

func isPrimo(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}
