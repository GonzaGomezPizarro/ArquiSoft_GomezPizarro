package div

import "testing"

func TestDivisionOK(t *testing.T) { //libreria testing (https://golang.org/pkg/testing/)
	Division, err := Division(10, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if Division != 5 {
		t.Error("Resultado incorrecto")
		return
	}
}
func TestDivisionErrorbyZero(t *testing.T) { //libreria testing (https://golang.org/pkg/testing/)
	_, err := Division(10, 0)
	if err == nil {
		t.Error("Error no esperado")
		return
	}
}
