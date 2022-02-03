package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var resultVictory int
var resultDefeat int

func main() {
	var wantToPlay string
	fmt.Println("Bem vindo ao Cara ou Coroa! Deseja jogar [sim/nao]?")
	fmt.Scan(&wantToPlay)
	switch strings.ToLower(wantToPlay) {
	case "sim", "s":
		play()
	case "nao", "n", "não":
		fmt.Println("Obrigado.")
	default:
		fmt.Println("Você precisa digitar sim ou nao! Execute o programa novamente.")
	}
}

func play() {
	remain := true
	for remain {
		var userCurrency string
		fmt.Println("Escolha cara ou coroa.")
		fmt.Scan(&userCurrency)
		if strings.ToLower(userCurrency) != "cara" && strings.ToLower(userCurrency) != "coroa" {
			fmt.Println("Por favor digite cara ou coroa.")
			continue
		}
		flipACoin(userCurrency)
		remain = playAgain()
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
	} else {
		fmt.Println(side, "Você Perdeu!")
		resultDefeat++
	}
}

func playAgain() bool {
	var again string
	fmt.Println("Deseja jogar novamente?")
	fmt.Scan(&again)
	switch strings.ToLower(again) {
	case "sim", "s":
		return true
	case "nao", "n", "não":
		fmt.Println("Ate a Proxima!", resultVictory, "vitoria", resultDefeat, "Derrota")
		return false
	default:
		fmt.Println("Por favor digite sim ou não! Por enquanto vamos continuar o jogo!")
		return true
	}
}
