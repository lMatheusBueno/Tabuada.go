package main

import "fmt"

func main() {
	fmt.Println("****Tabuada em Go!****")
	for i := 1; i <= 10; i++ { //primeiro laço fará 10 repetições
		for y := 1; y <= 10; y++ { //segundo laço fará 100 repetições

			fmt.Println(i, " x ", y, " = ", i*y)

		}
	}
}
