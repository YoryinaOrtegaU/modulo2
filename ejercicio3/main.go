package main

import "fmt"

func main() {
	meses := []string{"enero", "febrero", "octubre", "junio"}
	ObtenerEstacion(meses)

	ObtenerEstacion2(meses)

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

func ObtenerEstacion2(meses []string) {
	for _, mes := range meses {
		switch mes {
		case "enero", "febrero", "marzo":
			fmt.Println("Verano")
		case "abril", "mayo", "junio":
			fmt.Println("otoño")
		case "julio", "agosto", "septiembre":
			fmt.Println("invierno")
		case "octubre", "noviembre", "diciembre":
			fmt.Println("primavera")
		default:
			fmt.Println("no existe este mes")
		}
	}
}
