package main

import "fmt"

// if
func CalcularImpuestos(salario int) {
	if salario <= 3000 {
		fmt.Println("Debe pagar impuestos")
	} else if salario <= 4000 {
		impuesto := salario * 10 / 100
		fmt.Println("Debe pagar impuesto del 10%", impuesto)
	} else {
		impuesto := salario * 15 / 100
		fmt.Println("Debe pagar impuesto del 15%", impuesto)
	}
}

// switch sin condición
func VerificarEdad(edad int) {
	switch {
	case edad >= 100:
		fmt.Println("¿Eres inmortal?")
	case edad >= 18:
		fmt.Println("Eres mayor de edad")
	default:
		fmt.Println("Eres menor de edad")
	}
}

// Switch con multiples casos
func consultarDia(dia string) {
	switch dia {
	case "domingo", "Sabado":
		fmt.Println("Es un día del fin de semana")
	case "Lunes", "Martes", "Miercoles", "Jueves", "Viernes":
		fmt.Println("Es un día entre semana")
	default:
		fmt.Println("No es un día de la semana")
	}
}

// Switch con fallthrough
func verificarMultiploCinco() {
	for i := 0; i < 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("Fizz")
			fallthrough
		case i%5 == 0:
			fmt.Print("Buzz")
		case i%3 == 0:
			fmt.Print("Fizz2")
		default:
			fmt.Printf("%d", i)
		}
		fmt.Print(" ")
	}
}

func main() {
	//if
	var salario int
	salario = 3500
	CalcularImpuestos(salario)

	//Switch sin condicion
	var edad int = 10
	VerificarEdad(edad)

	// Switch con multiples casos
	var dia string
	dia = "Lunes"
	consultarDia(dia)

	// Switch con fallthrough
	verificarMultiploCinco()

	//for range
	frutas := []string{"mora", "pera", "mango"}

	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}

	//for "while"
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("sum", sum)

	//Continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "is odd")
	}

}
