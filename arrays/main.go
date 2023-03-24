package main

import "fmt"

// for range
func RecorrerSlice(array [3]string) {
	for i, value := range array {
		fmt.Println(i, value)
	}
}

func main() {
	// for range
	var frutas [3]string
	frutas[0] = "lulo"
	frutas[1] = "pera"
	frutas[2] = "mora"

	RecorrerSlice(frutas)

	var a [3]string
	a[0] = "hola"
	a[1] = "hora"
	a[2] = "ola"

	fmt.Println(a[0])
	fmt.Println(a)

	dias := [3]string{"lunes", "martes"}

	fmt.Println(dias, cap(dias), len(dias))
}
