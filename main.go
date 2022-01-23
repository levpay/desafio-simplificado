package main

import (
	"fmt"
	"math/rand"
)

var answer string
var victory int
var lost int

func main() {
	fmt.Println("Welcome to the game Heads or Tails!")       //bem vindo ao jogo Cara ou Coroa
	fmt.Print("would you like to play a game? 'y' or 'n': ") //Você gostaria de jogar uma partida 'y' or 'n'?

	fmt.Scan(&answer)
	//resposta
	switch answer {
	case "y":
		play()
	case "n":
		fmt.Println("Don't worry, if you change your mind, just click on the program again!.") //Tranquilo, caso mude de ideia , so clicar no programa novamente
	default:
		fmt.Println("incorrect option. try running the program again 'y' for yes or 'n' for no.") //opcao incorreta. tente executar o programa novamente  'y' para yes ou 'n' para no.
	}
}
func play() {
	play := true
	for play {
		fmt.Print("Which do you want heads or tails? [heads/tails] ") //Qual você quer cara ou coroa?
		fmt.Scan(&answer)
		in := []string{"heads", "tails"}
		randomIndex := rand.Intn(1)
		result := in[randomIndex]
		fmt.Println(result)
		if result == answer {
			fmt.Println("You got it right congratulations. It's at", result) //Você acertou parabéns. Está em
			victory++
		} else {
			fmt.Println("You were wrong unfortunately. It's at ", result) //Você errou infelizmente. Está em
			lost++
		}
		fmt.Print("Do you want to continue the game?'y' ou 'n' ") //Você quer continuar o jogo'y' ou 'n'?
		fmt.Scan(&answer)
		if answer == "n" {
			play = false
		}
	}
	fmt.Println("you ended the game!") //Você encerrou o jogo
	fmt.Println("your hits", victory)  //Seu acertos
	fmt.Println("your mistakes", lost) //Seus erros
}
