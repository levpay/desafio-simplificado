package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var resultVictory int
var resultDefeat int

func main() {
	var wantToPlay string
	fmt.Println("Bem vindo ao Cara ou Coroa! Deseja jogar [sim/nao]?")
	fmt.Scan(&wantToPlay)

	for {
		if strings.ToLower(wantToPlay) == "sim" || strings.ToLower(wantToPlay) == "s" {
			play()
		} else if strings.ToLower(wantToPlay) == "nao" || strings.ToLower(wantToPlay) == "n" {
			fmt.Println("Obrigado.")
			break
		}
		main()
		break
	}
}

func play() {
	var userCurrency string

	println("Escolha cara ou coroa.")
	fmt.Scan(&userCurrency)

	if strings.ToLower(userCurrency) == "cara" || strings.ToLower(userCurrency) == "coroa" {
		flipACoin(userCurrency)
	} else {
		fmt.Println("Por favor digite cara ou coroa.")
		play()
	}
}

func flipACoin(userCurrency string) {

	coin := []string{
		"cara",
		"coroa",
	}
	rand.Seed(time.Now().UnixNano())
	side := coin[rand.Intn(len(coin))]
	if strings.ToLower(userCurrency) == side {
		fmt.Println(side, "Você Ganhou!")
		resultVictory++
		playAgain()

	} else {
		fmt.Println(side, "Você Perdeu!")
		resultDefeat++
		playAgain()
	}
}

func playAgain() {
	var again string
	fmt.Println("Deseja jogar novamente?")
	fmt.Scan(&again)
	if strings.ToLower(again) == "sim" || strings.ToLower(again) == "s" {
		play()
	} else if strings.ToLower(again) == "nao" || strings.ToLower(again) == "n" {
		playResult(resultVictory, resultDefeat)
		fmt.Println("Ate a Proxima!")
		os.Exit(1)
	} else {
		fmt.Println("Por favor digite sim ou nao!")
	}
}

func playResult(resultVictory int, resultDefeat int) {
	fmt.Println(resultVictory, "vitoria", resultDefeat, "Derrota")
}
