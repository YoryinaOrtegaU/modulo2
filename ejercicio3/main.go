package main

import "fmt"

func main() {
	meses := []string{"enero", "febrero", "octubre", "junio"}
	ObtenerEstacion(meses)

	array1 := [2]int{1, 2}
	array2 := [2]string{"hola", "dos"}
	array3 := [2]string{"lola", "tres"}

	

	fmt.Println("array1", array1)
	fmt.Println("array2", array2)
	fmt.Println("referencia array2:", &array2)

	array2 = array3

	fmt.Println("array2", array2)
	fmt.Println("referencia array2:", &array2)

	var prueba []int
	var prueba2 [3]int
	prueba2[1] = 1


	fmt.Println("prueba", prueba, cap(prueba), len(prueba))
	fmt.Println("prueba2", prueba2, cap(prueba2), len(prueba2))

}

func ObtenerEstacion(meses []string) {
	for _, mes := range meses {
		if mes == "enero" || mes == "febrero" || mes == "marzo" {
			fmt.Println("Verano")
		} else if mes == "abril" || mes == "mayo" || mes == "junio" {
			fmt.Println("otoño")
		} else if mes == "julio" || mes == "agosto" || mes == "septiembre" {
			fmt.Println("invierno")
		} else if mes == "octubre" || mes == "noviembre" || mes == "diciembre" {
			fmt.Println("primavera")
		} else {
			fmt.Println("no existe este mes")
		}

	}
}
