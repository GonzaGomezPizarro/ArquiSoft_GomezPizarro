package div

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	a              float32
	b              float32
	ExpectedResult float32
	ExpectedError  bool
}

func TestDivisionOK(t *testing.T) { //libreria testing (https://golang.org/pkg/testing/)
	Divisio, err := Division(10, 2)

	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	assert.Nil(t, err) //libreria testing (https://golang.org/pkg/testing/)

	// if Division != 5 {
	// 	t.Error("Resultado incorrecto")
	// 	return
	// }
	assert.Equal(t, Divisio, float32(5)) // es lo mismo que el anterior

	for i := -2; i < 10; i++ {
		Divisio, err := Division(float32(i), 1.77)
		assert.Nil(t, err)
		assert.Equal(t, Divisio, float32(i)/1.77)
	}

	//Otra Forma:
	testcases := []TestCase{
		{a: 10, b: 2, ExpectedResult: 5, ExpectedError: false},
		{a: 10, b: 0, ExpectedResult: 999999, ExpectedError: true},
		{a: -10, b: 2, ExpectedResult: -5, ExpectedError: false},
		{a: -10, b: 0, ExpectedResult: 999999, ExpectedError: false}, // este esta mal para ver como muestra el error
		{a: 10, b: -2, ExpectedResult: -5, ExpectedError: false},
		{a: .5, b: 2, ExpectedResult: .24, ExpectedError: false}, // este esta mal para ver como muestra el error
	}
	for _, testcase := range testcases {
		Divisio, err := Division(testcase.a, testcase.b)
		assert.Equal(t, testcase.ExpectedResult, Divisio)
		assert.Equal(t, testcase.ExpectedError, err != nil)
	}

}
func TestDivisionErrorbyZero(t *testing.T) { //libreria testing (https://golang.org/pkg/testing/)
	_, err := Division(10, 0)

	// if err == nil {
	// 	t.Error("Error no esperado")
	// 	return
	// }
	assert.NotNil(t, err) //libreria testing (https://golang.org/pkg/testing/)

}
