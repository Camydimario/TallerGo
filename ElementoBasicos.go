//temperatura
package main

import "fmt"

func mostrarTemperatura() (float64, int) {
	const congelacion = 0
	var tempe float64 = 36.5
	fmt.Printf("la temperatura es de %f\n", tempe)
	fmt.Printf("La congelacion es de %d\n", congelacion)

	// TODO: declarar temperatura (float64 = 36.5) y constante Congelacion = 0,
	//       mostrarlos por pantalla con fmt.Println y devolverlos
	return tempe, congelacion
}

func main() {
	t, c := mostrarTemperatura()
	fmt.Println("los valores son", t, c)
}
//positivo y negativo
package main

import "fmt"

func clasificarNumero(n int) string {
	if n < 0 {
		fmt.Print("es negativo")
	} else if n == 0 {
		fmt.Print("es cero")
	} else {
		fmt.Print("es postivo")
	}
	// TODO: devolver "positivo", "negativo" o "cero" segun corresponda
	return "\n"
}

func main() {
	numero := clasificarNumero(-1)
	fmt.Printf(numero)

}

//dia de semana
package main

import "fmt"

func diaSemana(n int) string {

	switch n {
	case 1:
		return "es lunes"

	case 2:
		return "es martes"

	case 3:
		return "es miercoles"

	case 4:
		return "es jueves"

	case 5:
		return "es viernes"

	case 6:
		return "es sabado"

	case 7:
		return "es domingo"

	}

	// TODO: usando switch, devolver el dia de la semana (1 = "lunes" ... 7 = "domingo")
	return ""

}

func main() {
	diaDeLaSemana := diaSemana(7)
	fmt.Print(diaDeLaSemana)
}


//suma 
package main

import "fmt"

func sumarUnoACien() int {
	suma := 0
	for i := 0; i < 101; i++ {
		suma += i
	}

	// TODO: calcular la suma de los numeros del 1 al 100 usando un ciclo for
	return suma
}

func main() {
	sumador := sumarUnoACien()
	fmt.Print(sumador)
}


// maximo
package main

import "fmt"

// maximo encuentra y devuelve el valor máximo del slice usando range.
// Este mismo problema se retoma en el capítulo 07-punteros con la función
// Maximo(numeros []int, resultado *int), que escribe el resultado a través
// de un puntero en vez de devolverlo por retorno.
func maximo(numeros []int) int {
	maximo := numeros[0]
	for _, valor := range numeros {
		if valor > maximo {
			maximo = valor

		}

	}

	// TODO: encontrar y devolver el valor maximo del slice usando range
	return maximo
}

func main() {
	numeros := []int{10, 20, 30, 40, 50}
	fmt.Print("El conjunto de numeros es:\n", numeros)

	fmt.Print("\nel maximo es: \n", maximo(numeros))

}

