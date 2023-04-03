package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuesto(t *testing.T) {
	var salario float64 = 49000.0
	salario2 := 52000.0
	salario3 := 160000.0

	impuestoEsperado := 0.0
	impuestoEsperado2 := 8840.0
	impuestoEsperado3 := 43200.0

	impuesto := CalcularImpuesto(float64(salario))
	impuesto2 := CalcularImpuesto(float64(salario2))
	impuesto3 := CalcularImpuesto(float64(salario3))

	assert.Equal(t, impuestoEsperado, impuesto, "deben ser iguales")
	assert.Equal(t, impuestoEsperado2, impuesto2, "deben ser iguales")
	assert.Equal(t, impuestoEsperado3, impuesto3, "deben ser iguales")

}

func TestCalcularPromedio(t *testing.T) {
	promedioEsperado := 4

	promedio := Promedio(2, 3, 5, 5, 5, 4)

	assert.Equal(t, promedioEsperado, promedio, "deben ser iguales")
}

func TestCalcularSalario(t *testing.T) {
	categoria1 := "A"
	minutos := 15.0
	categoria2 := "B"
	minutos2 := 150.0
	categoria3 := "C"
	minutos3 := 50.0

	salarioEsperado1 := 975.0
	salarioEsperado2 := 4687.5
	salarioEsperado3 := 833.3333333333334

	salario1 := CalcularSalario(minutos, categoria1)
	salario2 := CalcularSalario(minutos2, categoria2)
	salario3 := CalcularSalario(minutos3, categoria3)

	assert.Equal(t, salarioEsperado1, salario1, "deben ser iguales")
	assert.Equal(t, salarioEsperado2, salario2, "deben ser iguales")
	assert.Equal(t, salarioEsperado3, salario3, "deben ser iguales")
}

func TestPromedio(t *testing.T) {
	resultadoEsperado := 3
	resultado := Promedio(1, 2, 3, 4, 5)
	assert.Equal(t, resultadoEsperado, resultado)
}

func TestCalcularMaximo(t *testing.T) {
	resultadoEsperado := 5
	resultado := CalcularMaximo(1, 2, 3, 4, 5)
	assert.Equal(t, resultadoEsperado, resultado)
}

func TestCalcularMinimo(t *testing.T) {
	resultadoEsperado := 1
	resultado := CalcularMinimo(1, 2, 3, 4, 5)
	assert.Equal(t, resultadoEsperado, resultado)
}

/* func TestStatistics(t *testing.T) {
    args := []int{10, 9, 8, 8, 4}
    t.Run(“Average”, func(t *testing.T) {
        expectedResult := 7.8
        result, err := averageFunc(args...)
        assert.Equal(t, expectedResult, result)
        assert.Nil(t, err)
    })
    t.Run(“Min”, func(t *testing.T) {
        expectedResult := 4.0
        result, err := minFunc(args...)
        assert.Equal(t, expectedResult, result)
        assert.Nil(t, err)
    })
    t.Run(“Max”, func(t *testing.T) {
        expectedResult := 10.0
        result, err := maxFunc(args...)
        assert.Equal(t, expectedResult, result)
        assert.Nil(t, err)
    })
} */

func TestAlimentoTara(t *testing.T) {
	numAnimales := 5
	resultadoEsperado := 750

	resultado := AlimentoTara(numAnimales)

	assert.Equal(t, resultadoEsperado, resultado)

}

func TestAlimentoHams(t *testing.T) {
	numAnimales := 5
	resultadoEsperado := 1250
	resultado := AlimentoHams(numAnimales)

	assert.Equal(t, resultadoEsperado, resultado)
}

func TestAlimentoPerr(t *testing.T) {
	numAnimales := 5
	resultadoEsperado := 50
	resultado := AlimentoPerr(numAnimales)

	assert.Equal(t, resultadoEsperado, resultado)
}

func TestAlimentoGat(t *testing.T) {
	numAnimales := 5
	resultadoEsperado := 25

	resultado := AlimentoGat(numAnimales)

	assert.Equal(t, resultadoEsperado, resultado)
}
