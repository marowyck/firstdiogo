// Você e seus colegas querem criar um código que exiba todos os números entre 1 e 100 que são divisíveis por 3.

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ { //mostra que o i é de 1 a 100
		if i%3 == 0 { // confere se é divisivel por 3
			fmt.Println(i) //printa o numero ne paizao
		}
	}
}
