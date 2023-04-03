package funciones

import "fmt"

func CalcularMinimo(notas ...int) int {
	notaMin := notas[0]
	for i := 0; i < len(notas); i++ {
		if notaMin > notas[i] {
			notaMin = notas[i]
		}
	}
	return notaMin
}

func CalcularMaximo(notas ...int) int {
	notaMax := notas[0]
	for i := 0; i < len(notas); i++ {
		if notaMax < notas[i] {
			notaMax = notas[i]
		}
	}
	return notaMax
}

func Promedio(notas ...int) int {
	var sumNotas int
	var promedio int
	for _, nota := range notas {
		sumNotas = sumNotas + nota
	}
	promedio = sumNotas / len(notas)
	return promedio
}

func CalcularEstadisticas(operacion string) func(notas ...int) int {
	switch operacion {
	case "minimo":
		return CalcularMinimo
	case "maximo":
		return CalcularMaximo
	case "promedio":
		return Promedio
	default:
		fmt.Println("La operación solicitada no existe")
		return nil
	}
}
