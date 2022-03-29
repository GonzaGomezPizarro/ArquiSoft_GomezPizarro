package main

import (
	"fmt"
	// importar el paquete div
)

func main() {

	var dividendo, divisor float32 = 10, 0

	div, err := division(dividendo, divisor)

	if err != nil {
		fmt.Println(" -> Error: ", err)
	}
	fmt.Println(" -> Resultado: ", div)

}
