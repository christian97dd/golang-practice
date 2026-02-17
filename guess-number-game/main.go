package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// fmt.Println(rand.Intn(100))
	play()
}

func play() {
	randomNum := rand.Intn(100)
	var enteredNum int
	var tries int
	const maxTries = 10

	for tries < maxTries {
		tries++
		fmt.Printf("Ingrese un numero (intentos restantes: %d):  ", maxTries-tries+1)
		fmt.Scanln(&enteredNum)

		if enteredNum == randomNum {
			fmt.Println("adivinaste el numero")
			playAgain()
			return
		} else if enteredNum < randomNum {
			fmt.Println("El numero es mayor")
		} else if enteredNum > randomNum {
			fmt.Println("El numero es menor")
		}
	}
	fmt.Println("no mas intentos, el numero era: ", randomNum)
	playAgain()
}

func playAgain() {
	var choose string
	fmt.Println("jugar de nuevo? (s/n): ")
	fmt.Scanln(&choose)

	switch choose {
	case "s":
		play()
	case "n":
		fmt.Println("vuelva pronto")
	default:
		fmt.Println("elegí s o n")
		playAgain()
	}
}
