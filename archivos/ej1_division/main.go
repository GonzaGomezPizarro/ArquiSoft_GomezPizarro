package main

import (
	"fmt"

	"github.com/GonzaGomezPizarro/ArquiSoft_GomezPizarro/archivos/ej1_division/div"
)

func main() {

	var dividendo, divisor float32 = 10, 0

	div, err := div.Division(dividendo, divisor)

	if err != nil {
		fmt.Println(" -> Error: ", err)
	}
	fmt.Println(" -> Resultado: ", div)

}
