package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var desejaJogar string
	var moeda string
	coin := []string{
		"cara",
		"coroa",
	}
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bem vindo ao Cara ou Coroa! Deseja jogar? ")
	fmt.Scan(&desejaJogar)
	if desejaJogar == "sim" {
		fmt.Println("Escolha! Cara ou Coroa?")
		fmt.Scan(&moeda)

	}
	side := coin[rand.Intn(len(coin))]

	if side == moeda {
		fmt.Println(side, "Voce Ganhou")

	} else {
		fmt.Println(side, "Voce perdeu!")
	}

}
