// promedio
package main

import "fmt"

func Promedio(nums ...float64) float64 {
	suma := 0.0
	cantidad := 0.0

	for _, n := range nums {
		suma += n
		cantidad++
	}
	return suma / cantidad
}

func main() {
	fmt.Print(Promedio(10, 4))
}


// aplicar funciones
package main

import "fmt"

func Aplicar(numeros []int, f func(int) int) []int {
	resultado := make([]int, len(numeros))

	for i, n := range numeros {
		resultado[i] = f(n)
	}
	return resultado
}
func sumarNumeros(n int) int {
	return n + 1
}
func main() {
	numeros := []int{10, 20, 30, 40, 50}
	resultadoFinal := Aplicar(numeros, sumarNumeros)
	fmt.Println(resultadoFinal)

}
