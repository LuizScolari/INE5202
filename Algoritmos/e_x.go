//algoritmo do eˆx
//quanto mais termos, mais aproximado será o resultado

package main

import (
	"fmt"
)

func main() {
	calculo := func(x float64, terms int) float64 {
		if terms < 0 {
			return 0
		}

		var somador float64 = 0
		fatorial := 1.0 // Começa com o fatorial de 0! (que é 1)
		potencia := 1.0 // Começa com x^0 (que é 1)

		for i := 0; i < terms; i++ {
			if i != 0 {
				fatorial *= float64(i) // Atualiza o fatorial acumulado
			}
			somador += potencia / fatorial
			potencia *= x // Atualiza x^i
		}
		return somador
	}
	resultado := calculo(2, 20)
	fmt.Println(resultado)
}
