package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Escolha entre Cara ou Coroa:") //Mostra as alternativas
	reader := bufio.NewReader(os.Stdin)         //Pede ao usuário a escolha
	input, err := reader.ReadString('\n')       // Le a entrada do usuário
	if err != nil {
		fmt.Println(err) // Se der erro, vai imprimi-lo
	}
	fmt.Println(input) // Se não der erro imprimi

}
