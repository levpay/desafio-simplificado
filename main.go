package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var wantToPlay string

	fmt.Println("Bem vindo ao Cara ou Coroa! Deseja jogar? ")
	fmt.Scan(&wantToPlay)
	if strings.ToLower(wantToPlay) == "sim" || strings.ToLower(wantToPlay) == "s" {
		play()
	} else if strings.ToLower(wantToPlay) == "nao" || strings.ToLower(wantToPlay) == "n" {
		fmt.Println("Obrigado")
	} else {
		fmt.Println("Por favor digite sim ou nao!")
	}

}

func play() {
	var userCurrency string

	println("Escolha cara ou coroa")
	fmt.Scan(&userCurrency)

	if strings.ToLower(userCurrency) == "cara" || strings.ToLower(userCurrency) == "coroa" {
		flipACoin(userCurrency)
	} else {
		fmt.Println("Por favor digite cara ou coroa")
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
		fmt.Println(side, "Você Ganhou")
		playAgain()

	} else {
		fmt.Println(side, "Você Perdeu!")
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
		fmt.Println("Ate a Proxima!")
	} else {
		fmt.Println("Por favor digite sim ou nao!")
	}

}
