package defer1

import (
	"fmt"
	"os"
)

func LeerArchivo2(archivo string) {
	_, err := os.Open(archivo)
	defer fmt.Println("ejecución finalizada")
	fmt.Println("antes de paniquear")
	if err != nil {
		panic("El archivo indicado no fue encontrado o esta daniado")
	}
}
