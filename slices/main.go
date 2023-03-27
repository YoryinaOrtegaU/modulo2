package main

import "fmt"

func main() {

	//slice
	var sl []int
	sl = append(sl, 4, 1, 2)
	sl[2] = 3

	fmt.Println("sl", sl, cap(sl), len(sl))

	var ola = []int{2, 3, 4}
	ola = append(ola, 3, 6, 7, 8, 9)
	fmt.Println("ola", ola, ola[:3], cap(ola), len(ola))
	fmt.Println("posicion 1", ola[1])

	nombres := make([]string, 5)
	nombres = append(nombres, "esperancito")
	fmt.Println("nombres", nombres, cap(nombres), len(nombres))
}
