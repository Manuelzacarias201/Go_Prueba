package main

import "fmt"

func main() {

	//hola mundo
	fmt.Println("Hello world")
	/*
	   esto es un comentario
	*/

	//variables
	var myString string = "esto es una cadena de texto"
	fmt.Println(myString)

	myString = "aqui cambio el valor de la cadena de texto"
	fmt.Println(myString)

	var myString2 = "Esto es una cadena de texto"
	fmt.Println(myString2)

	//myString = 6 error por que se esta refiriendo a un estring y nosotros estamos usando un numero de tipo int

	//variable de tipo int
	var numDos int = 2025
	fmt.Println(numDos)
}
