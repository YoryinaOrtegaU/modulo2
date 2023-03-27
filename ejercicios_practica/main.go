package main

import "fmt"

func imprimirNumeroLetrasPalabra(palabra string) {

	var numeroLetras int

	for i := 0; i < len(palabra); i++ {
		numeroLetras += 1
	}

	fmt.Println("La palabra", palabra, "tiene: ", numeroLetras, "letras")
}

type Cliente struct {
	edad       int
	empleado   bool
	antiguedad int
	sueldo     int
}

func VerificarClientePrestamos(cliente Cliente) {

	if cliente.edad > 22 && cliente.empleado == true && cliente.antiguedad > 1 {
		fmt.Println("El cliente cumple con los requisitos de edad, empleabilidad y antiguedad en el cargo")
		if cliente.sueldo > 100000 {
			fmt.Println("No se le cobrará interés ")
		}
	} else {
		fmt.Println("El cliente no cumple con los requisitos")

	}
	/*
		 	switch {
			case cliente.edad > 22:
				fmt.Println("El cliente cumple con la edad")
			case cliente.empleado:
				fmt.Println("cumple con ser empleado")
			case cliente.antiguedad > 1:
				fmt.Println("El cliente cumple con más de un año")
			case cliente.sueldo > 100000:
				fmt.Println("No se le cobrará interés ")
			}
	*/
}

func ImprimirMesCorrespondiente(numMes int) {

	if numMes < 13 {
		var meses = map[int]string{
			1:  "Enero",
			2:  "Febrero",
			3:  "Marzo",
			4:  "Abril",
			5:  "Mayo",
			6:  "Junio",
			7:  "Julio",
			8:  "Agosto",
			9:  "Septiembre",
			10: "Octubre",
			11: "Noviembre",
			12: "Diciembre",
		}
		mes := meses[numMes]

		fmt.Println("El mes correspondiente al número:", numMes, "es", mes)
	} else {
		fmt.Println("El número no corresponde a un mes")
	}
}

func ImprimirMesCorrespondiente2(numMes int) {

	switch numMes < 13 {
	case numMes == 1:
		fmt.Println("El mes correspondiente al número:", numMes, "es Enero")
	case numMes == 2:
		fmt.Println("El mes correspondiente al número:", numMes, "es Febrero")
	case numMes == 3:
		fmt.Println("El mes correspondiente al número:", numMes, "es Marzo")
	case numMes == 4:
		fmt.Println("El mes correspondiente al número:", numMes, "es Abril")
	case numMes == 5:
		fmt.Println("El mes correspondiente al número:", numMes, "es Mayo")
	case numMes == 6:
		fmt.Println("El mes correspondiente al número:", numMes, "es Junio")
	case numMes == 7:
		fmt.Println("El mes correspondiente al número:", numMes, "es Julio")
	case numMes == 8:
		fmt.Println("El mes correspondiente al número:", numMes, "es Agosto")
	case numMes == 9:
		fmt.Println("El mes correspondiente al número:", numMes, "es Septiembre")
	case numMes == 10:
		fmt.Println("El mes correspondiente al número:", numMes, "es Octubre")
	case numMes == 11:
		fmt.Println("El mes correspondiente al número:", numMes, "es Noviembre")
	case numMes == 12:
		fmt.Println("El mes correspondiente al número:", numMes, "es Diciembre")
	}
}

func main() {
	var palabra string
	palabra = "hola"

	numeroLetras := len(palabra)
	fmt.Println(numeroLetras)

	imprimirNumeroLetrasPalabra(palabra)

	cliente1 := Cliente{
		edad:       23,
		empleado:   true,
		antiguedad: 2,
		sueldo:     105000,
	}

	VerificarClientePrestamos(cliente1)

	ImprimirMesCorrespondiente(13)
	ImprimirMesCorrespondiente(12)
}
