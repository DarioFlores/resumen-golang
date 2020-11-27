package main

import "fmt"

var empleados map[string]int

func main() {
	empleados = make(map[string]int)

	empleados["Grupo Esfera"] = 50
	empleados["Mercado Libre"] = 12000

	fmt.Println(empleados)
	fmt.Println(empleados["Grupo Esfera"])
}
