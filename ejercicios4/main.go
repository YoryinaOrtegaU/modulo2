package main

import "fmt"

func main() {

	Counter()

	var item = [4]int{1, 2, 3, 4}
	IterateArray(item)

	var frutas = []string{"banana", "manzana", "pera", "uva", "banana", "manzana", "pera", "uva", "banana", "manzana"}

	IterateSlice(frutas)

	var registro = make(map[string][]string)
	registro["juan"] = []string{"banana", "manzana", "pera", "uva"}
	registro["pello"] = []string{"naranja", "fresa", "mora"}
	registro["mirta"] = []string{"mango", "mora", "guineo"}
	fmt.Println(registro)
	delete(registro, "juan")
	fmt.Println(registro)

	//verificar si existe un item en el mapa
	items, ok := registro["pello"]
	if ok {
		fmt.Println("- items de pello: ", items)
	} else {
		fmt.Println("- items de juan, no encontrado")
	}

	for persona, lista := range registro {
		fmt.Println("persona:", persona)
		fmt.Println("frutas:", lista)
	}

	//IterateMap()
	Iterate(frutas)
}

func Counter() {
	// Crear un incrementador de a 2
	// - solo se debe repitir 5 veces
	var lim int = 5
	var num int = 10

	for i := 0; i < lim; i++ {
		num += 2
		fmt.Println(num)
	}
}

func IterateArray(array [4]int) {
	// Crear un array de items
	// - mostrar el primer item
	// - mostrar el último item
	fmt.Println("primer item", array[0])
	fmt.Println("último item", array[len(array)-1])

}

func IterateSlice(sli []string) {
	// Crear un slice de items
	// - mostrar el primer item
	// - mostrar el último item
	// - agregar un item al slice
	// - verificar si el item existe

	fmt.Println("primer item", sli[0])
	fmt.Println("último item", sli[len(sli)-1])
	sli = append(sli, "papaya")
	fmt.Println("slice con elemento agregado", sli)

}

func IterateMap(map[string][]string) {
	// Crear un registro de personas y sus items
	// - mostrar la persona y sus items
	// - agregar una persona y sus items
	// - mostrar el registro de todas las personas y sus items

}

func Iterate(frutas []string) {
	// A partir de un slice de items, contar cuántas veces se repite cada item y mostrarlo en un registro
	var mapaContador = make(map[string]int)

	for _, fruta := range frutas {
		if _, ok := mapaContador[fruta]; ok {
			mapaContador[fruta] += 1
		} else {
			mapaContador[fruta] = 1
		}
	}

	fmt.Println(mapaContador)
}
