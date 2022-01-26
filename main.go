package main

import ( //Biblioteca
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Escolha entre Cara ou Coroa:") //Mostra as alternativas
	fmt.Println("Cara: 0")
	fmt.Println("Coroa: 1")
	reader := bufio.NewReader(os.Stdin)   //Pede ao usuário a escolha
	input, err := reader.ReadString('\n') // Lê a entrada do usuário
	if err != nil {
		fmt.Println(err) // Se der erro, vai imprimi-lo
	}
	input = strings.TrimSuffix(input, "\n") //Remove o \n do final da string
	opcao, err := strconv.Atoi(input)       // converte a string para inteiro

	if opcao > 1 { //Se o usuário digitar uma opção inválida
		fmt.Println("Opção Inválida") //Imprimi a mensagem
		return
	}
	status, result := jogo(opcao) // Chama a função jogo
	if status == 1 {
		fmt.Println("Você ganhou!")
		fmt.Println("O resultado foi: ", result)
	} else {
		fmt.Println("Você perdeu!")
		fmt.Println("O resultado foi: ", result)
	}
}
func jogo(opcao int) (int, string) { //Retorna o Status do jogo e o resultado
	rand.Seed(time.Now().UnixNano())     //Gera um número aleatório
	options := []string{"cara", "coroa"} // Opções possíveis
	result := rand.Intn(2)               //Gera um número aleatório entre 0 e 1

	if options[result] == options[opcao] {
		return 1, options[result] //Retorna 1 se o usuário ganhou
	} else {
		return 2, options[result] //Retorna 2 se o usuário perdeu
	}
}
