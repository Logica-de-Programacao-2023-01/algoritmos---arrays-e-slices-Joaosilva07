package main

import "fmt"

func main() {
    var matriz [2][3]int

    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("Informe o valor do elemento [%d][%d]: ", i, j)
            fmt.Scanln(&matriz[i][j])
        }
    }

    var linha, coluna int
    fmt.Print("Informe o índice de linha: ")
    fmt.Scanln(&linha)
    fmt.Print("Informe o índice de coluna: ")
    fmt.Scanln(&coluna)

    valor := matriz[linha][coluna]
    fmt.Println("Valor na posição:", valor)
}
