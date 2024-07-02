// Você e seus colegas querem criar em formato de código uma antiga brincadeira:
// Ao falar os números sempre que aparecer um múltiplo de 3 o participante deve falar "pin"
// e nos múltiplos de 5 deve falar "pan"
// vamo lá, fé!

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ { //confere que é de 1 a 100
		switch { //switch case nessa ocasião é god
		case i%3 == 0 && i%5 == 0:
			fmt.Println("PinPan")
		case i%3 == 0:
			fmt.Println("Pin")
		case i%5 == 0:
			fmt.Println("Pan")
		default:
			fmt.Println(i)
		}
	}
}
