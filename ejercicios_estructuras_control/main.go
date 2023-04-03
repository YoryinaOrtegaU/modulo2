package main

import "fmt"

// Ejercicio numero de letras por palabra
func imprimirNumeroLetrasPalabra(palabra string) {

	var numeroLetras int

	for i := 0; i < len(palabra); i++ {
		numeroLetras += 1
	}

	fmt.Println("La palabra", palabra, "tiene: ", numeroLetras, "letras")
}

// Ejercicio prestamos
type Cliente struct {
	edad       int
	empleado   bool
	antiguedad int
	sueldo     int
}

func VerificarClientePrestamos(cliente Cliente) {
	VerificarEdadCliente(cliente)
	VerificarClienteEmpleado(cliente)
	VerificarAntiguedeadCliente(cliente)
	VerificarAplicaInteresCliente(cliente)
}

func VerificarEdadCliente(cliente Cliente) {
	if cliente.edad > 22 {
		fmt.Println("El cliente cumple con la edad")
	} else {
		fmt.Println("El cliente no cumple con la edad")
	}
}

func VerificarClienteEmpleado(cliente Cliente) {
	if cliente.empleado == true {
		fmt.Println("El cliente cumple con ser empleado")
	} else {
		fmt.Println("El cliente no cumple con el requisito de ser empleado")
	}
}

func VerificarAntiguedeadCliente(cliente Cliente) {
	if cliente.antiguedad > 1 {
		fmt.Println("El cliente cumple con la antiguead en el cargo requerida")
	} else {
		fmt.Println("El cliente no cumple con el requisito de la antiguead en el cargo ")
	}
}

func VerificarAplicaInteresCliente(cliente Cliente) {
	if cliente.sueldo > 100000 {
		fmt.Println("No se le cobrará interés ")
	} else {
		fmt.Println("Al cliente se le cobrará intereses")
	}
}

// Ejercicios imprimir mes correspondiente con if
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

// Ejercicios imprimir mes correspondiente con switch
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

func contarEmpleados(employees map[string]int) int {
	var contador int
	for _, edad := range employees {
		if edad > 21 {
			contador++
		}
	}
	return contador
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

	ImprimirMesCorrespondiente(13)
	ImprimirMesCorrespondiente(12)

	VerificarClientePrestamos(cliente1)

	//Ejercicio que
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Bejamin es:", employees["Benjamin"])
	contadorEmpleados := contarEmpleados(employees)
	fmt.Println(contadorEmpleados)
	employees["Federico"] = 25

	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)
}
