package main

import "fmt"

//Um professor de ensino médio solicitou aos seus alunos que desenvolvessem um programa para converter o valor do ponto de ebulição da água de K para °C//

var ebulicaoK = 373.2

func main() {
	tempK := ebulicaoK
	tempC := (tempK - 100)
	fmt.Printf("A temperatura de ebulição em K é de: %g, enquanto a temperatura em °C é de: %g", tempK, tempC)
}
