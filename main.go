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

var vitorias int = 0
var derrotas int = 0
var continuar bool = true

func main() {
	for continuar == true { //Inicia o loop
		fmt.Println("Escolha entre Cara ou Coroa:") //Mostra as alternativas
		fmt.Println("Cara: 0")
		fmt.Println("Coroa: 1")

		reader := bufio.NewReader(os.Stdin) //Pede ao usuário a escolha

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
			vitorias++
			fmt.Println("O resultado foi: ", result)
		} else {
			fmt.Println("Você perdeu!")
			derrotas++
			fmt.Println("O resultado foi: ", result)
		}
		fmt.Println("Deseja continuar jogando? (s/n)")
		reader = bufio.NewReader((os.Stdin))                      //Pede ao usuário a entrada de sim ou não
		input_continuar, err_continuar := reader.ReadString('\n') //Lê a entrada do usuário
		if err_continuar != nil {
			fmt.Println(err_continuar) //Se der erro
		}
		input = strings.TrimSuffix(input_continuar, "\n") //Remove o \n do final da string
		if input == "s" {                                 //se o usuário digitar s
			continuar = true //Continua o jogo
		} else {
			continuar = false                                // se não, para o jogo
			fmt.Println("Obrigado por jogar, volte sempre!") // Frase de despedida caso o usuário não vá mais jogar
		}
	}
	fmt.Println("Vitórias: ", vitorias) // Depois que o usuário parar de jogar mostra o resultado de vitórias/derrotas
	fmt.Println("Derrotas: ", derrotas)
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
