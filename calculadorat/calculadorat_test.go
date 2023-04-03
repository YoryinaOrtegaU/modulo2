package calculadorat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	a := 1
	b := 3
	resultadoEsperado := 4

	resultado := Sumar(a, b)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
