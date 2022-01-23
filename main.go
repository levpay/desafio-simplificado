package main

import (
	"fmt"
)

var answer string

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
	fmt.Print("")
}
