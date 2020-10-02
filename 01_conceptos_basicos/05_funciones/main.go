package main

import (
	"fmt"
	"strings"
)

func main() {

	var a, b int = 8, 3

	suma := sumar(a, b)
	resta := restar(a, b)

	fmt.Printf("Suma:\n %d + %d = %d\n", a, b, suma)
	fmt.Printf("Resta:\n %d - %d = %d\n", a, b, resta)

	imprimirMensajes(74, 1, "hola", -5)

	s := "CaDeNa De EjEmPlO"

	minusculas, mayusculas, titulo := transformarCadena(s)

	fmt.Printf("Transformación de cadena de caracteres:\n"+
		" Cadena original: %s\n Cadena en minúsculas: %s\n"+
		" Cadena en mayúsculas: %s\n Cadena como título: %s\n",
		s, minusculas, mayusculas, titulo)

}

func sumar(x, y int) int {
	return x + y
}

func restar(x int, y int) (resta int) {
	resta = x - y
	return
}

func imprimirMensajes(m, n int, s string, o int) {
	fmt.Printf("La función imprimirMensajes no tiene salida\n")
	fmt.Printf("Tiene tres entradas numéricas:\n m: %d\n n: %d\n o: %d\n", m, n, o)
	fmt.Printf("Y tiene una entrada de cadena de caracteres:\n s: %s\n", s)
}

func transformarCadena(s string) (string, string, string) {
	minusculas := strings.ToLower(s)
	mayusculas := strings.ToUpper(s)
	titulo := strings.Title(strings.ToLower(s))

	return minusculas, mayusculas, titulo
}
