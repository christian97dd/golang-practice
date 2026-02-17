package main

import (
	"fmt"
	"strconv"
)

func main() {
	// integers: int o unit, la diferencia unit solo almacena positivos
	// var integer uint
	// var float float32

	// fmt.Println(math.MaxFloat32)

	fullName := "Christian de Diego \t(alias \"Cdd\")\n"

	fmt.Println(fullName)

	// bytes:
	var a byte = 'b'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])

	// rune:
	var r rune = '🩵'
	fmt.Println(r)

	// valores default
	var (
		defaulInt    int
		defaultUint  uint
		defaulFloat  float32
		defaulBool   bool
		defaulString string
	)

	fmt.Println(defaulInt, defaultUint, defaulFloat, defaulBool, defaulString)

	// convercion de variables
	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(integer16 + int16(integer32))

	// convertir de string a number con strconv.Atoi
	stringToNumber := "100"
	convertedStringInNumber, error := strconv.Atoi(stringToNumber)

	fmt.Println(convertedStringInNumber, error)

	// convertir de number a string con strconv.Itoa
	numero := 42
	numberToString := strconv.Itoa(numero)
	fmt.Println("numberToString", numberToString)

	// name := "Christian"
	// age := 28
	var name string
	var age int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola, me llamo %s y tengo %d años. \n", name, age)

	// greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age)
	// fmt.Println(greeting)

	fmt.Printf("el tipo de name es: %T\n", name)
	fmt.Printf("el tipo de age es: %T\n", age)
}
