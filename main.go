package main

import (
	"fmt"
	"math/rand"
)

var answer string
var vitoria int
var perdeu int

func main() {
	fmt.Println("Welcome to the game Heads or Tails!")  //bem vindo ao jogo Cara ou Coroa
	fmt.Print("would you like to play a game? [y/n]: ") //Você gostaria de jogar uma partida?

	fmt.Scan(&answer)
	//resposta
	switch answer {
	case "y":
		play()
	case "n":
		fmt.Println("Don't worry, if you change your mind, just click on the program again.") //Tranquilo, caso mude de ideia , so clicar no programa novamente
	default:
		fmt.Println("incorrect option. try running the program again 'y' for yes or 'n' for no.") //opcao incorreta. tente executar o programa novamente  'y' para yes ou 'n' para no.
	}
}
func play() {
	play := true
	for play {
		fmt.Println("Which do you want heads[h] or tails[t]?") //Qual você quer cara ou coroa?
		fmt.Scan(&answer)
		in := []string{"heads", "tails"}
		randomIndex := rand.Intn(len(in))
		result := in[randomIndex]
		if result == answer {
			fmt.Println("Você acertou. caiu em ", result)
			vitoria++
		} else {
			fmt.Println("Você errou. caiu em ", result)
			perdeu++
		}
		fmt.Println("Você deseja continuar o jogo?[y] ou [n]")
		fmt.Scan(&answer)
		if answer == "n" {
			play = false
		}
	}
	fmt.Println("Você encerrou o jogo!")
	fmt.Println("Seu acertos", vitoria)
	fmt.Println("Seus erros", perdeu)
}
