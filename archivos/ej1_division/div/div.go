package div

import "errors"

//si la primer letra es mayuscula el metodo es publico fuera del paquete

func Division(a, b float32) (float32, error) {
	if b == 0 {
		return 999999, errors.New("No se puede dividir por cero")
	}
	var res float32 = a / b
	return res, nil
}

func DivInt(a, b int) (int, int) {
	return a / b, a % b
}

// en este caso las funciones son publicas fuera del paquete
