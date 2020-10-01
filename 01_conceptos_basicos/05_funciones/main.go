package main

import "fmt"

func main() {

	var a, b int = 8, 3

	suma := sumar(a, b)
	resta := restar(a, b)

	fmt.Printf("Suma:\n %d + %d = %d\n", a, b, suma)
	fmt.Printf("Resta:\n %d - %d = %d\n", a, b, resta)

	imprimirMensajes(74, 1, "hola", -5)
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
