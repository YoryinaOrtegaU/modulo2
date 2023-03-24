package main

import "fmt"

func prueba(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func main() {

	var (
		temperatura = 31
		humedad     = 49
		presion     = 100
	)

	fmt.Println("temperatura: ", temperatura)
	fmt.Println("humedad: ", humedad)
	fmt.Println("presion:", presion)

	fmt.Println(prueba('p'))
}
