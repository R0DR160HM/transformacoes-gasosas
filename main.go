package main

import (
	"fmt"
	"strconv"
	"strings"

)

func main() {

	greetings()

	for {
		constantTemperature := false
		constantPressure := false
		constantBulk := false

		fmt.Println("A temperatura é constante? [S]im / [N]ão")
		constantTemperature = getAnswer("a temperatura não é constante")

		if !constantTemperature {
			fmt.Println("A pressão é constante? [S]im / [N]ão")
			constantPressure = getAnswer("a pressão não é constante")

			if !constantPressure {
				fmt.Println("O volume é constante? [S]im / [N]ão")
				constantBulk = getAnswer("o volume não é constante")
			}
		}

		details(constantBulk, constantPressure, constantTemperature)

		fmt.Println("Deseja continuar? [S]im / [N]ão")
		ok := getAnswer("você deseja continuar")

		if !ok {
			fmt.Println("Até mais")
			break
		}
	}
}

func greetings() {
	fmt.Println("Feito por Rodrigo Heinzen de Moraes")
	fmt.Println("https://github.com/R0DR160HM/transformacoes-gasosas")
}

func getAnswer(statement string) bool {
	var command string
	fmt.Scan(&command)

	switch strings.ToUpper(command) {
	case "S":
		return true
	case "N":
		return false
	default:
		fmt.Println("Comando", command, "não reconhecido (esperado S/N). Assumiremos que", statement)
		return false
	}
}

func getValue(sentence string) float32 {
	fmt.Println(sentence)
	var command string
	fmt.Scan(&command)

	num, err := strconv.ParseFloat(command, 32)

	if err != nil {
		fmt.Println("Valor", command, "não reconhecido.")
		return getValue(sentence)
	}

	return float32(num)
}

func details(constantBulk bool, constantPressure bool, constantTemperature bool) {
	var initialBulk float32
	var finalBulk float32
	var initialPressure float32
	var finalPressure float32
	var initialTemperate float32
	var finalTemperature float32

	if constantBulk {
		initialBulk = 1
		finalBulk = 1
	} else {
		initialBulk = getValue("Insira o volume inicial do gás:")
		finalBulk = getValue("Insira o volume final do gás:")
	}

	fmt.Println(initialBulk, finalBulk, initialPressure, finalPressure, initialTemperate, finalTemperature)
}
