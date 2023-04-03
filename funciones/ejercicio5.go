package funciones

import "fmt"

func AlimentoTara(numAnimales int) int {
	cantidadAlimentoTara := numAnimales * 150
	return cantidadAlimentoTara
}

func AlimentoHams(numAnimales int) int {
	cantidadAlimentoHams := numAnimales * 250
	return cantidadAlimentoHams
}

func AlimentoPerr(numAnimales int) int {
	cantidadAlimentoPerr := numAnimales * 10
	return cantidadAlimentoPerr
}

func AlimentoGat(numAnimales int) int {
	cantidadAlimentoGat := numAnimales * 5
	return cantidadAlimentoGat
}

func Animal(animal string) func(numAnimales int) int {
	switch animal {
	case "tarantula":
		return AlimentoTara
	case "hamster":
		return AlimentoHams
	case "perro":
		return AlimentoPerr
	case "gato":
		return AlimentoGat
	default:
		fmt.Println("No contamos con ese animal")
		return nil
	}
}
