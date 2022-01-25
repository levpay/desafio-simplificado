package main

import ( //Biblioteca
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Escolha entre Cara ou Coroa:") //Mostra as alternativas
	reader := bufio.NewReader(os.Stdin)         //Pede ao usuário a escolha
	input, err := reader.ReadString('\n')       // Lê a entrada do usuário
	if err != nil {
		fmt.Println(err) // Se der erro, vai imprimi-lo
	}
	fmt.Println(input) // Se não der erro imprimi

}
func jogo(opcao int) (int, string) { //Retorna o Status do jogo e o resultado
	rand.Seed(time.Now().UnixNano())                  //Gera um número aleatório
	options := []string{"cara", "coroa"}              // Opções possíveis
	result := rand.Intn(2)                            //Gera um número aleatório entre 0 e 1
	fmt.Println("O resultado foi: ", result)          //Imprime o resultado
	fmt.Println("O resultado foi: ", options[result]) //Imprime o resultado

	if options[result] == options[opcao] {
		return 1, options[result] //Retorna 1 se o usuário ganhou
	} else {
		return 2, options[result] //Retorna 2 se o usuário perdeu
	}
}
