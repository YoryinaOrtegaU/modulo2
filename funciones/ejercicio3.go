package funciones

import "fmt"

func CalcularSalario(minutos float64, categoria string) float64 {
	horasTrabajadas := minutos / 60
	var salario float64

	switch categoria {
	case "C":
		salario = 1000 * horasTrabajadas
	case "B":
		salario = 1500 * horasTrabajadas
		salario = salario + (salario * 25 / 100)
	case "A":
		salario = 3000 * horasTrabajadas
		salario = salario + (salario * 30 / 100)
	default:
		fmt.Println("Categoria no encontrada")
	}
	return salario
}
