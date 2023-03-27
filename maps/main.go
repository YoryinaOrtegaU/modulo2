package main

import "fmt"

func main() {

	//Puedo agregar los datos en las llaves
	mapa := map[string]int{
		"Julio": 2,
		"Juan":  3,
	}

	mapa["lolo"] = 23
	fmt.Println("mapa", mapa["Juan"])
	fmt.Println("mapa", mapa)

	mapa["Juan"] = 40

	fmt.Println("mapa", mapa["Juan"])

	delete(mapa, "Juan")

	fmt.Println("mapa", mapa)

	//Inicializa un mapa, pero los datos no se pueden agrefar enseguida
	mapa2 := make(map[string]int)
	mapa2["tote"] = 19
	mapa2["perencejo"] = 30

	fmt.Println(mapa["tote"])

	fmt.Println(mapa)
}
