package main

import (
	"fmt"
	"go/types"
)

const a = 10

func main() {
	fmt.Println("****Tabuada em Go!****")
	for i := 1; i <= 10; i++ { //primeiro laço fará 10 repetições
		for y := 1; y <= 10; y++ { //segundo laço fará 100 repetições

			fmt.Println(i, " x ", y, " = ", i*y)

		}
		fmt.Print("_______________\n") //quebra linha
	}
}
